package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Setup sets up model for gin
func Setup(db *gorm.DB) {

	log.Println("----< AutoMigrating [if not migrated] >-----")

	// Automigrate schema
	db.AutoMigrate(&Workflow{})

}
