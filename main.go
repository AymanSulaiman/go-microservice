package main

import (
	"go-microservice/config"
	"go-microservice/database"
	"go-microservice/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	database.Conn()

	r := gin.Default()

	r.GET("/", handlers.Homepage)

	// blog post test
	r.GET("/test", handlers.TestPage)

	// blog get post
	r.GET("/posts/:id", handlers.GetPost)

	// blog make post
	r.POST("/posts", handlers.CreatePost)

	// running server
	r.Run(":8080")
}
