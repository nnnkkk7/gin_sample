package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/message", func(c *gin.Context) {
		message := c.PostForm("message")
		nickname := c.DefaultPostForm("nick", "none")

		c.JSON(200, gin.H{
			"status":   "ok",
			"message":  message,
			"nickname": nickname,
		})
	})

	router.Run(":8080")
}
