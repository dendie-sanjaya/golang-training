package main

import (
	"fmt"
	"log"
	"net"
	grpcHandler "praisindo/handler/grpc"
	pb "praisindo/proto/user_service/v1"
	"praisindo/repository/postgres_gorm"
	"praisindo/service"

	"praisindo/config"
	"praisindo/entity"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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
	err = gormDB.AutoMigrate(&entity.User{})
	if err != nil {
		fmt.Println("Failed to migrate database schema user saldo:", err)
	} else {
		fmt.Println("Database schema migrated user saldo successfully")
	}

	// setup gorm connection
	//dsn := "postgresql://postgres:12345678@localhost:5432/Golang"
	//dsn = config.PostgresStringConnection
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	// setup service

	// uncomment to use postgres gorm
	userRepo := postgres_gorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	//userHandler := ginHandler.NewUserHandler(userService)
	userHandler := grpcHandler.NewUserHandler(userService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50051")
	_ = grpcServer.Serve(lis)
}
