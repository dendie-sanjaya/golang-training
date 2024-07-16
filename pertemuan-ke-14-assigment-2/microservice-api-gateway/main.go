package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "praisindo/proto"
)

func main() {
	conn, err := grpc.NewClient(
		"0.0.0.0:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterUserWalletServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := gin.Default()
	gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	log.Println("Running grpc gateway server in port :8080")
	_ = gwServer.Run()
}
