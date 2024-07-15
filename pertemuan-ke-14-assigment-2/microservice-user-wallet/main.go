package main

import (
	"fmt"
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

	//untuk topup
	//rtn, _ := postgres_gorm.TopUp(gormDB, 1, 1000000)
	//rtn, _ = postgres_gorm.TopUp(gormDB, 2, 500000)

	//untuk transfer
	rtn, _, _ := postgres_gorm.Transfer(gormDB, 1, 2, 2000000)
	fmt.Println(rtn)

	//untuk get user balance
	// saldo, _ := postgres_gorm.GetUserBalance(gormDB, 2)
	// str := fmt.Sprintf("%.2f", saldo)
	// fmt.Println("Saldo user 2 : Rp. " + str)

	//untuk get history
	// rtn, _ := postgres_gorm.GetTransactionHistory(gormDB, 1)
	// fmt.Println(rtn)

	// setup service
	// uncomment to use postgres gorm
	//userRepo := postgres_gorm.NewUserWalletH(gormDB)

	//fmt.Println(userRepo)
	//userService := service.NewUserService(userRepo)
	//userHandler := ginHandler.NewUserHandler(userService)
	//userHandler := grpcHandler.NewUserHandler(userService)

	userWalletHandler := postgres_gorm.NewUserWalletHandler(gormDB)
	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserWalletServiceServer(grpcServer, userWalletHandler)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		log.Println("Running grpc server in port :50052")
		_ = grpcServer.Serve(lis)
	}()
	//time.Sleep(1 * time.Second)

	// go func() {
	// 	log.Println("Running grpc server in port :50051")
	// 	_ = grpcServer.Serve(lis)
	// }()
	// time.Sleep(1 * time.Second)

	// // Run the grpc gateway
	// conn, err := grpc.NewClient(
	// 	"0.0.0.0:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalln("Failed to dial server:", err)
	// }
	// gwmux := runtime.NewServeMux()
	// if err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// // dengan default http server
	// /*
	// 	gwServer := &http.Server{
	// 		Addr:    ":8080",
	// 		Handler: gwmux,
	// 	}
	// 	log.Println("Running grpc gateway server in port :8080")
	// 	_ = gwServer.ListenAndServe()
	// */

	// // dengan GIN
	// gwServer := gin.Default()
	// gwServer.Group("v1/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	// log.Println("Running grpc gateway server in port :8080")
	// _ = gwServer.Run()
}
