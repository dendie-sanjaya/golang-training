package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"praisindo/config"
	"praisindo/entity"
	pb "praisindo/proto"
	postgres_gorm "praisindo/repository/postgres_gorm"
	"time"

	"google.golang.org/grpc"
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
	err = gormDB.AutoMigrate(&entity.UserSaldo{})
	if err != nil {
		fmt.Println("Failed to migrate database schema user saldo:", err)
	} else {
		fmt.Println("Database schema migrated user saldo successfully")
	}

	// Migrate the schema
	err = gormDB.AutoMigrate(&entity.UserSaldoHistory{})
	if err != nil {
		fmt.Println("Failed to migrate database schema user_saldo:", err)
	} else {
		fmt.Println("Database schema migrated user_saldo_history successfully")
	}

	err = gormDB.AutoMigrate(&entity.UserWalet{})
	if err != nil {
		fmt.Println("Failed to migrate database schema user walet :", err)
	} else {
		fmt.Println("Database schema migrated user walet successfully")
	}

	//setup redis connection
	/*
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.RedisHost,
			Password: config.RedisPassword, // no password set
			DB:       config.RedisDatabase, // use default DB
		})
	*/

	userWalletHandler := postgres_gorm.NewUserWalletHandler(gormDB)

	// Create a new context
	ctx := context.Background()
	//userWalletHandler.CreateWallet(ctx)
	//userWalletHandler.UpdateWallet(ctx)
	//userWalletHandler.DeleteWallet(ctx)
	//userWalletHandler.GetUserBalanceByWallet(ctx)
	userWalletHandler.Spend(ctx)

	grpcServer := grpc.NewServer()
	pb.RegisterUserWalletServiceServer(grpcServer, userWalletHandler)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50052")
	_ = grpcServer.Serve(lis)

}
