package routes

import "github.com/gin-gonic/gin"

// SetupRoutes sets up the routes
func SetupRoutes(routes *gin.Engine) {
	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
