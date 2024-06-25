package main

import (
	"log"
	"praisindo/handler"
	"praisindo/repository/postgres_gorm"
	"praisindo/router"
	"praisindo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	dsn := "postgresql://postgres:12345678@localhost:5432/Golang_Assigment"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := postgres_gorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	submissionRepo := postgres_gorm.NewsubmissionRespository(gormDB)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)

	// // Routes
	router.SetupRouter(r, userHandler, submissionHandler)
	//router.SetupRouter(r, userHandler)

	//Run the server
	log.Println("Running server on port 8181")
	r.Run(":8181")
}
