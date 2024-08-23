package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"praisindo_consumer_1/config"
	"praisindo_consumer_1/repository/postgres_gorm"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ============== HELPER FUNCTIONS ==============

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/notifications", func(ctx *gin.Context) {
		get_data(ctx.Writer, ctx.Request)
	})

	// Jalankan server pada port 8080
	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func get_data(w http.ResponseWriter, r *http.Request) {
	// Inisialisasi koneksi database
	dsn := "host=" + config.PostgresHost + " port=" + config.PostgressPort + " user=" + config.PostgresUser + " password=" + config.PostgresPassword + " dbname= " + config.PostgresDB + " sslmode=" + config.PostgresSSLMode
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}

	// Inisialisasi UserHandler
	userHandler := postgres_gorm.UserHandler{Db: gormDB}

	// Ambil semua notifikasi dari database
	notifications, err := userHandler.GetAllNotifications(context.Background())
	if err != nil {
		http.Error(w, "Failed to retrieve notifications", http.StatusInternalServerError)
		return
	}

	// Konversi data ke format JSON
	jsonData, err := json.Marshal(notifications)
	if err != nil {
		http.Error(w, "Failed to convert notifications to JSON", http.StatusInternalServerError)
		return
	}

	// Set header dan tulis data JSON ke response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
