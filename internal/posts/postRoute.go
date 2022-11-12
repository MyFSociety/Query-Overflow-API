package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "list of all posts",
	})
}

func getPostById(c *gin.Context) {
	postId := c.Param("id")
	message := "returned post id with " + postId
	c.JSON(http.StatusOK, gin.H{
		"data": message,
	})
}

func createPost(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"data": "post created",
	})
}

func updatePost(c *gin.Context) {
	postId := c.Param("id")
	message := "updated post id with " + postId
	c.JSON(http.StatusOK, gin.H{
		"data": message,
	})
}

func deletePost(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}

func Routes(route *gin.Engine) {
	posts := route.Group("/posts")
	{
		posts.GET("/", getAllPosts)
		posts.GET("/:id", getPostById)
		posts.POST("/create", createPost)
		posts.PATCH("/:id", updatePost)
		posts.DELETE("/:id", deletePost)
	}
}
