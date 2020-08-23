package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dipeshdulal/workflow/models"
	"github.com/dipeshdulal/workflow/routes"
	"github.com/dipeshdulal/workflow/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file. %v", err)
	}

	app := gin.Default()

	db := utils.DBSetup(app)
	defer db.Close()

	models.Setup(db)
	app.Use(func(c *gin.Context) {
		fmt.Println("Middleware called.")
		c.Set("db", db)
	})

	log.Println("---------------------------")
	log.Println("----- WORKFLOW ENGINE -----")
	log.Println("---------------------------")

	routes.SetupRoutes(app)
	app.Run(":" + os.Getenv("APP_PORT"))
}
