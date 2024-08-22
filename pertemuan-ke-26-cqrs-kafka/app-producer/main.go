package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"praisindo_producer/entity"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

// ============== HELPER FUNCTIONS ==============
func findUserByID(id int) (entity.User, error) {
	for _, user := range entity.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, entity.ErrUserNotFoundInProducer
}

// ============== KAFKA RELATED FUNCTIONS ==============
func sendKafkaMessage(producer sarama.SyncProducer, ctx *gin.Context, fromID, toID int) error {
	message := ctx.PostForm("message")

	fromUser, err := findUserByID(fromID)
	if err != nil {
		return err
	}

	toUser, err := findUserByID(toID)
	if err != nil {
		return err
	}

	notification := entity.Notification{
		From: fromUser,
		To:   toUser, Message: message,
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: entity.DefaultTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJSON),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}

func sendMessageHandler(producer sarama.SyncProducer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fromID, err := strconv.Atoi(ctx.PostForm("fromID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		toID, err := strconv.Atoi(ctx.PostForm("toID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = sendKafkaMessage(producer, ctx, fromID, toID)
		if errors.Is(err, entity.ErrUserNotFoundInProducer) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully!",
		})
	}
}

func setupProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{entity.KafkaBrokerAddress},
		config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}

func main() {
	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/send", sendMessageHandler(producer))

	fmt.Printf("Kafka PRODUCER started at http://localhost%s\n",
		entity.KafkaProducerPort)

	if err := router.Run(entity.KafkaProducerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
