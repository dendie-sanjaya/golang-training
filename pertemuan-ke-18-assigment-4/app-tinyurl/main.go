package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"praisindo/config"
	grpcHandler "praisindo/handler/grpc"
	pb "praisindo/proto/shorturl_service/v1"
	postgres_gorm "praisindo/repository/postgres_gorm"
	"praisindo/service"
	"time"

	"github.com/gin-gonic/gin"

	"praisindo/entity"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	time.Sleep(3 * time.Second)
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
	err = gormDB.AutoMigrate(&entity.ShortUrl{})
	if err != nil {
		fmt.Println("Failed to migrate database schema:", err)
	} else {
		fmt.Println("Database schema migrated successfully")
	}

	//setup redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword, // no password set
		DB:       config.RedisDatabase, // use default DB
	})

	shortUrlRepo := postgres_gorm.NewshortUrlRepository(gormDB)
	shortUrlService := service.NewShortUrlService(shortUrlRepo, rdb)

	// rst, _ := shortUrlService.CreateShortUrl("https://detik.com/id/u/mengkonversi/detik/ke/bulan#7776000")
	// fmt.Println(rst)
	// rst2, _ := shortUrlService.GetShortUrl("dWgXDZ")
	// fmt.Println(rst2)

	//userHandler := ginHandler.NewUserHandler(userService)
	shortUrlHandler := grpcHandler.NewShortUrlHandler(shortUrlService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterShortUrlServiceServer(grpcServer, shortUrlHandler)

	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Running grpc server in port :50052")
		_ = grpcServer.Serve(lis)
	}()

	time.Sleep(1 * time.Second)

	// Run the grpc gateway
	conn, err := grpc.NewClient(
		"0.0.0.0:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterShortUrlServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := gin.Default()

	// Definisikan redirect ke url asli
	gwServer.GET("/v1/getshorturl/:url_short", func(c *gin.Context) {
		urlTarget := c.Param("url_short")
		urlShort, err := shortUrlService.GetShortUrl(urlTarget)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		log.Print(urlShort.UrlLong)
		c.Redirect(http.StatusMovedPermanently, urlShort.UrlLong)
	})

	gwServer.Run(":8080")
}
