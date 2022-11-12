package main

import (
	"net/http"

	post "harry/query-overflow/internal/posts"

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

	post.Routes(router)

	router.Run("localhost:8080")

}
