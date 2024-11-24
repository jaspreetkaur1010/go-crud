package main

import (
	"github.com/jaspreetkaur1010/go-crud/initializers"
	"github.com/jaspreetkaur1010/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}