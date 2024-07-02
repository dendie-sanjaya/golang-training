package main

import (
	"context"
	"log"
	"net"

	pb "praisindo/proto"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.
	pb.Result
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello2(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello ke 2: " + in.Name}, nil
}

func main() {
	runServer()
}

func runServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
