package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"praisindo_consumer_1/config"
	"praisindo_consumer_1/entity"
	"praisindo_consumer_1/repository/postgres_gorm"
	"strconv"
	"sync"

	"golang.org/x/exp/slog"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ============== HELPER FUNCTIONS ==============
var ErrNoMessagesFound = errors.New("no messages found")

func getUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrNoMessagesFound
	}
	return userID, nil
}

// ====== NOTIFICATION STORAGE ======
type UserNotifications map[string][]entity.Notification

type NotificationStore struct {
	data UserNotifications
	mu   sync.RWMutex
}

func (ns *NotificationStore) Add(userID string, notification entity.Notification) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.data[userID] = append(ns.data[userID], notification)
}

func (ns *NotificationStore) Get(userID string) []entity.Notification {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.data[userID]
}

// ============== KAFKA RELATED FUNCTIONS ==============
type Consumer struct {
	store       *NotificationStore
	userHandler *UserNotifications // Tambahkan field userHandler
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)
		var notification entity.Notification
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}
		slog.Info("Consuming notification and adding it to storage", slog.Any("notification", notification))
		consumer.store.Add(userID, notification)
		db_save(&notification)

		sess.MarkMessage(msg, "")
	}

	return nil
}

func setupConsumerGroup(ctx context.Context, store *NotificationStore) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup([]string{entity.KafkaBrokerAddress}, entity.DefaultGroupID, config)
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		store: store,
	}

	for {
		err = consumerGroup.Consume(ctx, []string{entity.DefaultTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func handleNotifications(ctx *gin.Context, store *NotificationStore) {
	userID, err := getUserIDFromRequest(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	notes := store.Get(userID)
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message":       "No notifications found for user",
				"notifications": []entity.Notification{},
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}

func main() {
	db_migrate()

	store := &NotificationStore{
		data: make(UserNotifications),
	}

	ctx, cancel := context.WithCancel(context.Background())
	go setupConsumerGroup(ctx, store)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		handleNotifications(ctx, store)
	})

	fmt.Printf("Kafka CONSUMER (Group: %s) 👥📥 "+
		"started at http://localhost%s\n", entity.DefaultGroupID, entity.KafkaConsumerPort)

	if err := router.Run(entity.KafkaConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}

func db_migrate() {
	// Setup gorm connection without selecting a database
	dsn := "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " sslmode=" + config.PostgresSSLMode
	//dsn := "postgresql://postgres:password@postgres:5434/postgres?sslmode=disable"
	fmt.Println(dsn)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database connection established", config.PostgresHost, config.PostgressPort, config.PostgresSSLMode)
	}

	// Check if the database exists
	var exists bool
	err = gormDB.Raw("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = ?)", config.PostgresDB).Scan(&exists).Error
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database exists:", exists, config.PostgresDB)
	}

	// Create the database if it does not exist
	if !exists {
		err = gormDB.Exec("CREATE DATABASE " + config.PostgresDB).Error
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("Database created successfully")
		}
	}
	// Reconnect to the newly created database
	dsn = "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname= " + config.PostgresDB + " sslmode=" + config.PostgresSSLMode
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	err = gormDB.AutoMigrate(&entity.NotificationLog{})
	if err != nil {
		fmt.Println("Failed to migrate database schema notification log", err)
	} else {
		fmt.Println("Database schema migrated user")
	}
}

func db_save(NotificationLog *entity.Notification) {
	// Inisialisasi koneksi database
	dsn := "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname= " + config.PostgresDB + " sslmode=" + config.PostgresSSLMode
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	// // Inisialisasi UserHandler
	userHandler := postgres_gorm.UserHandler{Db: gormDB}

	// Contoh notifikasi yang akan disimpan
	logEntry := &entity.NotificationLog{
		FromId:  strconv.Itoa(NotificationLog.From.ID),
		ToId:    strconv.Itoa(NotificationLog.To.ID),
		Message: NotificationLog.Message,
	}

	// // Simpan notifikasi ke database
	if err := userHandler.SaveNotification(context.Background(), logEntry); err != nil {
		log.Printf("failed to save notification log: %v", err)
	} else {
		log.Println("notification log saved successfully")
	}
}
