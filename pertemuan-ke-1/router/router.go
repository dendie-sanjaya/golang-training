package router

import (
	"praisindo/handler"
	"praisindo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	//r.Use(gin.Logger())
	r.GET("/", handler.RootHandler)
	r.POST("/post", handler.PostHandler)

	//Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndpoint := r.Group("/private")

	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}

}
