package main

import (
	"log"
	"net"
	pb "praisindo/proto"
	"praisindo/repository/postgres_gorm"

	grpc "google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// setup gorm connection
	dsn := "postgresql://postgres:12345678@localhost:5432/Golang_Wallet"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	userWalletHandler := postgres_gorm.NewUserWalletHandler(gormDB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserWalletServiceServer(grpcServer, userWalletHandler)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50052")
	_ = grpcServer.Serve(lis)
}
