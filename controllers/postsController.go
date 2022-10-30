package controllers

import (
	"go-api/initializers"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(ctx *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Post created successfully",
		"post": post,
	})
}

func FindAllPost(ctx *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, posts)
}

func FindPostById(ctx *gin.Context) {
	var post models.Post
	id := ctx.Param("id")
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, post)
}

func UpdatePostById(ctx *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	id := ctx.Param("id")
	// find the post
	var post models.Post
	initializers.DB.First(&post, id)

	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	ctx.JSON(200, gin.H{
		"message":     "Post updated successfully",
		"updatedData": post,
	})
}

func DeletePostById(ctx *gin.Context) {
	var post models.Post
	id := ctx.Param("id")
	initializers.DB.Delete(&post, id)
	ctx.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
