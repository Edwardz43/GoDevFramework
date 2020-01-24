package routes

import (
	"Edwardz43/godevframework/controllers"

	"github.com/gin-gonic/gin"
)

// Serve ...
func Serve() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", controllers.Hello)

	SetAPI(r)

	// r.Run(":8765") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r
}
