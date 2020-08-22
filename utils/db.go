package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DBSetup is called to setup database
func DBSetup(app *gin.Engine) *gorm.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
	)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatalf("couldn't connect to database. %v", err)
	}
	return db
}
