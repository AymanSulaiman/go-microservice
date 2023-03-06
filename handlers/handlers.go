// handlers/handlers.go

package handlers

import (
	"strconv"

	"go-microservice/controller"

	"github.com/gin-gonic/gin"
)

// Homepage
func Homepage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to my blog!",
	})
}

// Test Page
func TestPage(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "Test",
	})
}

// Get a post

func GetPost(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid post id"})
		return
	}

	post, err := controller.GetPost(id)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"could": "Not get from database"})
		return
	}
	ctx.JSON(200, post)
}

func CreatePost(ctx *gin.Context) {
	var post controller.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid post creation"})
	}
	id, err := controller.CreatePost(&post)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Error creating post"})
	}
	ctx.JSON(201, gin.H{"id": id})
}
