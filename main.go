package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/bootstrap"
	"go-crud/controllers"
)

func init() {

	bootstrap.LoadEnv()
	bootstrap.ConnectDatabase()

}

func main() {

	r := gin.Default()
	r.POST("/post", controllers.PostStore)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.GET("/post/:id", controllers.PostShow)
	r.GET("/posts", controllers.PostIndex)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()

}
