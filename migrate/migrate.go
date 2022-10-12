package main

import (
	"go-crud/bootstrap"
	"go-crud/models"
	"fmt"
)

func init() {
	
	bootstrap.LoadEnv()
	bootstrap.ConnectDatabase()
}

func main() {
	bootstrap.DB.AutoMigrate(&models.Post{})
	fmt.Println("Migration completed")


}