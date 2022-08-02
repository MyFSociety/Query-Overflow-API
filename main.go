package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All Good :)",
		})
	})

	// get all posts
	router.GET("/posts", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"data": "list of all posts",
		})
	})

	// get a post
	router.GET("/posts/:id", func(c *gin.Context) {
		postId := c.Param("id")
		message := "returned post id with " + postId
		c.JSON(http.StatusOK, gin.H{
			"data": message,
		})
	})

	// create a post
	router.POST("/posts/create", func(c *gin.Context) {

		c.JSON(http.StatusCreated, gin.H{
			"data": "post created",
		})
	})

	// update a post
	router.PATCH("/posts/:id", func(c *gin.Context) {
		postId := c.Param("id")
		message := "updated post id with " + postId
		c.JSON(http.StatusOK, gin.H{
			"data": message,
		})
	})

	// delete a post
	router.DELETE("/posts/:id", func(c *gin.Context) {

		c.JSON(http.StatusNoContent, gin.H{})
	})

	router.Run("localhost:8080")

}
