package controllers

import "github.com/gin-gonic/gin"

// TestGet as a controller for the test purpose.
func TestGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get test",
	})
}

// TestPost ...
func TestPost(c *gin.Context) {
	message := c.DefaultPostForm("message", "message not found")
	nick := c.DefaultPostForm("nickname", "nickname not found")

	c.JSON(200, gin.H{
		"status":   "posted",
		"message":  message,
		"nickname": nick,
	})
}
