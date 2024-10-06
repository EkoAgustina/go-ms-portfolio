// Package aboutmodels defines the data structures for "About" entities.
package aboutmodels

import (
	"gorm.io/gorm"
)

// About represents an "About" entity in the database.
// It includes fields for storing information related to the "About" section of a portfolio.
//
// Fields:
// - ID: Auto-generated ID for the about entry (inherited from gorm.Model).
// - CreatedAt: Timestamp for when the about entry was created (inherited from gorm.Model).
// - UpdatedAt: Timestamp for when the about entry was last updated (inherited from gorm.Model).
// - Content: The content of the "About" section, marked as required for validation.
type About struct {
	gorm.Model
	Content string `json:"content" binding:"required"` // Content of the "About" section
}
