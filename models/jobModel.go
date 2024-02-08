package models

import "gorm.io/gorm"

// Job model for Gorm
type Job struct {
	gorm.Model // Includes fields ID, CreatedAt, UpdatedAt, DeletedAt
	Role       string
	Company    string
	Status     string
	Notes      string
}
