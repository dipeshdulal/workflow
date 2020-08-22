package main

import (
	"fmt"
	"log"
	"os"

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
	routes.SetupRoutes(app)

	db := utils.DBSetup(app)
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	defer db.Close()

	fmt.Println("---------------------------")
	fmt.Println("----- WORKFLOW ENGINE -----")
	fmt.Println("---------------------------")

	app.Run(":" + os.Getenv("APP_PORT"))
}
