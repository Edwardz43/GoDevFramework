package controllers

import (
	"Edwardz43/godevframework/services/api"

	"github.com/gin-gonic/gin"
)

var (
	apiservice *api.Service
)

func init() {
	apiservice = api.GetInstance()
}

// CreateUser ...
func CreateUser(c *gin.Context) {

	name := c.PostForm("name")

	email := c.PostForm("email")

	if name == "" || email == "" {
		c.JSON(200, gin.H{
			"status":    "posted",
			"isSuccess": false,
		})
	} else {
		result := apiservice.CreateNewUser(name, email)

		c.JSON(200, gin.H{
			"status":    "posted",
			"isSuccess": result,
		})
	}

}

// FindUserByID ...
func FindUserByID(c *gin.Context) {

	id := c.PostForm("id")

	if id == "" {
		c.JSON(200, gin.H{
			"status":    "posted",
			"isSuccess": false,
		})
	} else {
		result := apiservice.FindUserByID(id)

		c.JSON(200, gin.H{
			"status": "posted",
			"user":   result,
		})
	}

}
