package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaspreetkaur1010/go-crud/controllers"
	"github.com/jaspreetkaur1010/go-crud/initializers"
)
func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}