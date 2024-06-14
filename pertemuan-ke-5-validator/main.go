package main

import (
	"log"
	"praisindo/entity"
	"praisindo/handler"
	"praisindo/repository/slice"
	"praisindo/router"
	"praisindo/service"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// // setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// // Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8181")
	r.Run(":8181")
}
