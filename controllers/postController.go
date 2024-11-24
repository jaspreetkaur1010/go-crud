package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaspreetkaur1010/go-crud/initializers"
	"github.com/jaspreetkaur1010/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	//get data off req body
	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)


	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context){
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
    // Get the ID from the URL
    id := c.Param("id")
    intID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid ID format"})
        return
    }

    // Fetch the post by ID
    var post models.Post
    if err := initializers.DB.First(&post, intID).Error; err != nil {
        c.JSON(404, gin.H{"error": "Post not found"})
        return
    }

    // Return the post
    c.JSON(200, gin.H{
        "post": post,
    })
}

func PostsUpdate(c *gin.Context){
	//Get the id of the URL
	id:= c.Param("id")

	//Get the data off the req body
	var body struct{
		Title string
		Body string
	}

	c.Bind(&body)

	//Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)


	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})

	//Respond it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context){
	id:= c.Param("id")

	initializers.DB.Delete(&models.Post{},id)

	c.Status(200)
}

