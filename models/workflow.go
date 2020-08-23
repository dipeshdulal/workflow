package models

import "github.com/jinzhu/gorm"

// Workflow model
type Workflow struct {
	gorm.Model
	Type      string `json:"type"`
	SearchKey string `json:"search_key"`
}
