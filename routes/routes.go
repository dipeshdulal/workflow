package routes

import (
	"github.com/dipeshdulal/workflow/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes
func SetupRoutes(routes *gin.Engine) {
	api := routes.Group("/api/v1")
	{
		api.POST("/workflow", controllers.SaveWorkflow)
	}
}
