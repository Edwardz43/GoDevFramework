package routes

import (
	"Edwardz43/godevframework/controllers"

	"github.com/gin-gonic/gin"
)

// SetAPI ...
func SetAPI(r *gin.Engine) {
	api := r.Group("/api")
	{
		test := api.Group("/test")
		{
			test.GET("/get", controllers.TestGet)
			test.POST("/post", controllers.TestPost)
		}
		user := api.Group("/user")
		{
			user.POST("/create", controllers.CreateUser)
			user.POST("/findone", controllers.FindUserByID)
		}
	}
}
