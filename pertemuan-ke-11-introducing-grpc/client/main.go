package main

import (
	"context"
	"log"
	pb "praisindo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	runClient()
}

func runClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	greeterClient := pb.NewGreeterClient(conn)

	name := "world"
	r, err := greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
	log.Println()

	name = "dendie"
	r2, err2 := greeterClient.SayHello2(context.Background(), &pb.HelloRequest{Name: name})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err)
		log.Println()
	}

	log.Printf("Greeting: %s", r2.GetMessage())
	log.Println()

}
