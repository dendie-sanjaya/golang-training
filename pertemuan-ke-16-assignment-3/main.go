package main

import (
	"fmt"
	"log"
	"praisindo/config"
	postgres_gorm "praisindo/repository/postgres_gorm"
	"praisindo/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// setup gorm connection
	dsn := "postgresql://postgres:12345678@localhost:5432/Golang_shorturl"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	//setup redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword, // no password set
		DB:       config.RedisDatabase, // use default DB
	})

	shortUrlRepo := postgres_gorm.NewshortUrlRepository(gormDB)
	shortUrlService := service.NewShortUrlService(shortUrlRepo, rdb)

	rst, _ := shortUrlService.CreateShortUrl("https://detik.com/id/u/mengkonversi/detik/ke/bulan#7776000")
	fmt.Println(rst)
	rst2, _ := shortUrlService.GetShortUrl("dWgXDZ")
	fmt.Println(rst2)

	//userHandler := ginHandler.NewUserHandler(userService)
	//userHandler := grpcHandler.NewUserHandler(userService)

	// Run the grpc server
	// grpcServer := grpc.NewServer()
	// pb.RegisterUserServiceServer(grpcServer, userHandler)
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// log.Println("Running grpc server in port :50051")
	// _ = grpcServer.Serve(lis)
}
