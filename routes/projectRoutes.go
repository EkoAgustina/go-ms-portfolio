package routes

import (
	"github.com/EkoAgustina/go-ms-portfolio/controllers/projectControllers"
	"github.com/EkoAgustina/go-ms-portfolio/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupProjectRoutes configures routes for "project" endpoints on the Gin router.
// This function sets up the following routes:
// - POST /addProject: Creates a new "project" entity. Validated with ValidateApiKey middleware.
// - GET /project: Retrieves a "project" entity by ID or all projects. Validated with ValidateApiKey middleware.
//
// Parameters:
// - router: The Gin router instance to configure.
//
// Example:
//   router := gin.Default()
//   routes.SetupProjectRoutes(router)
func SetupProjectRoutes(router *gin.Engine) {
	router.POST("/addProject", middlewares.ValidateApiKey(), projectcontrollers.CreateProject)
	router.GET("/project", middlewares.ValidateApiKey(), projectcontrollers.GetProject)
}
