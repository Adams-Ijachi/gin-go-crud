package controllers

import (
	"go-crud/bootstrap"
	"go-crud/forms"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostStore(c *gin.Context) {

	var post models.Post

	if err := c.ShouldBind(&post); err != nil {
		_, errorMessages := forms.ValidateForm(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
			"errors":  errorMessages,
		})
		return
	}

	post = models.Post{Name: post.Name, Body: post.Body}

	result := bootstrap.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(500, gin.H{"message": "Error occured", "error": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"message": "sing",
		"data":    post,
	})
}

func PostUpdate(c *gin.Context) {

	postId := c.Param("id")

	if postId == "" {
		c.JSON(400, gin.H{"message": "Post id is required"})
		return
	}

	var body models.Post
	var post models.Post

	if err := c.ShouldBind(&body); err != nil {
		_, errorMessages := forms.ValidateForm(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
			"errors":  errorMessages,
		})
		return
	}

	bootstrap.DB.First(&post, postId)

	bootstrap.DB.Model(&post).Updates(body)

	c.JSON(200, gin.H{
		"message": "sing",
		"data":    post,
	})
}

func PostDelete(c *gin.Context) {

	postId := c.Param("id")

	if postId == "" {
		c.JSON(400, gin.H{"message": "Post id is required"})
		return
	}

	var post models.Post

	bootstrap.DB.First(&post, postId)

	if post.ID == 0 {
		c.JSON(404, gin.H{"message": "Post not found"})
		return
	}

	bootstrap.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "sing",
		"data":    post,
	})
}

// show all posts

func PostIndex(c *gin.Context) {

	var posts []models.Post

	bootstrap.DB.Find(&posts)

	c.JSON(200, gin.H{
		"message": "sing",
		"data":    posts,
	})
}

// show single post

func PostShow(c *gin.Context) {

	postId := c.Param("id")

	if postId == "" {
		c.JSON(400, gin.H{"message": "Post id is required"})
		return
	}

	var post models.Post

	bootstrap.DB.First(&post, postId)

	c.JSON(200, gin.H{
		"message": "sing",
		"data":    post,
	})
}
