package contactmodels

import (
	"gorm.io/gorm"
)

// Contact represents a contact entity in the database.
// It includes fields for storing information about a contact message from a user.
//
// Fields:
// - ID: Auto-generated ID for the contact (inherited from gorm.Model).
// - CreatedAt: Timestamp for when the contact was created (inherited from gorm.Model).
// - UpdatedAt: Timestamp for when the contact was last updated (inherited from gorm.Model).
// - Name: The name of the person who contacted.
// - Email: The email address of the person who contacted.
// - Subject: The subject of the contact message.
// - Message: The content of the contact message.
type Contact struct {
	gorm.Model
	Name    string `json:"name"`    // Name of the person who contacted
	Email   string `json:"email"`   // Email address of the person who contacted
	Subject string `json:"subject"` // Subject of the contact message
	Message string `json:"message"` // Content of the contact message
}
