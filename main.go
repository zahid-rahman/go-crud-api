package main

import (
	"go-api/controllers"
	"go-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	server := gin.Default()

	var posts = server.Group("/api/posts")
	posts.POST("/", controllers.PostCreate)
	posts.GET("/", controllers.FindAllPost)
	posts.GET("/:id", controllers.FindPostById)
	posts.PATCH("/:id", controllers.UpdatePostById)
	posts.DELETE("/:id", controllers.DeletePostById)

	server.Run()
}
