package main

import (
	"context"
	"log"
	"praisindo/handler"
	"praisindo/repository/postgres_pgx"
	"praisindo/router"
	"praisindo/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pgxPool, err := connectDB("postgresql://postgres:12345678@localhost:5432/Golang")
	if err != nil {
		log.Fatalln(err)
	}

	// setup service

	// slice db is disabled. uncomment to enabled
	//var mockUserDBInSlice []entity.User
	//_ = slice.NewUserRepository(mockUserDBInSlice)

	// pgx db is enabled. comment to disabled
	userRepo := postgres_pgx.NewUserRepository(pgxPool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8181")
	r.Run(":8181")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
