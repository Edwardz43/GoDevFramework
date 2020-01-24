package controllers

import "github.com/gin-gonic/gin"

// Hello return hello world as demo message.
func Hello(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"message": "hello world",
		},
	)
}
