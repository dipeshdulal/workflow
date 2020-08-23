package controllers

import (
	"log"

	"github.com/dipeshdulal/workflow/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type postWorkflow struct {
	ID string `json:"id"`
}

// SaveWorkflow saves the workflow to database
func SaveWorkflow(c *gin.Context) {
	workflow := postWorkflow{}
	if err := c.ShouldBindJSON(&workflow); err != nil {
		c.JSON(500, gin.H{
			"message": "couldn't bind json",
		})
		log.Printf("validation error: %v \n", err)
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	workflowDB := models.Workflow{SearchKey: workflow.ID, Type: "fsm"}
	if err := db.Create(&workflowDB).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error.",
		})
		log.Printf("database error: %v \n", err)
		return
	}

	c.JSON(200, gin.H{
		"message": "workflow saved",
		"data": gin.H{
			"id":         workflowDB.ID,
			"search_key": workflowDB.SearchKey,
			"type":       workflowDB.Type,
		},
	})
}
