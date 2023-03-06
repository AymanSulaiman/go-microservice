package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// homepage
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "This is a blog, do you like it?",
		})
	})

	// blog post test
	r.GET("/posts/:postID", func(ctx *gin.Context) {
		postID := ctx.Param("postID")
		ctx.JSON(200, gin.H{
			"msg": "Requested post " + postID,
		})
	})

	// running server
	r.Run(":8080")
}
