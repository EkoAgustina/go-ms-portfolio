package projectmodels

import (
	"gorm.io/gorm"
)

// Project represents a project entity in the database.
// It includes fields for storing project details such as title, description, image, and repository link.
//
// Fields:
// - ID: Auto-generated ID for the project (inherited from gorm.Model).
// - CreatedAt: Timestamp for when the project was created (inherited from gorm.Model).
// - UpdatedAt: Timestamp for when the project was last updated (inherited from gorm.Model).
// - ImageTitle: The title of the project's image.
// - Image: The URL or path to the project's image.
// - ProjectTitle: The title of the project.
// - ProjectDescription: A brief description of the project.
// - RepositoryLink: A link to the project's repository (e.g., GitHub).
type Project struct {
	gorm.Model
	ImageTitle       string `json:"imageTitle"`       // Title of the project's image
	Image            string `json:"image"`            // URL or path to the project's image
	ProjectTitle     string `json:"projectTitle"`     // Title of the project
	ProjectDescription string `json:"projectDescription"` // Description of the project
	RepositoryLink   string `json:"repositoryLink"`   // Link to the project's repository
}
