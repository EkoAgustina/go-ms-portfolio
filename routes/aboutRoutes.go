
package routes

import (
	"github.com/EkoAgustina/go-ms-portfolio/controllers/aboutControllers"
	"github.com/EkoAgustina/go-ms-portfolio/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupAboutRoutes configures routes for "about" endpoints on the Gin router.
// This function sets up the following routes:
// - POST /createAbout: Creates a new "about" entity. Validated with ValidateApiKey middleware.
// - GET /about: Retrieves an "about" entity by ID. Validated with ValidateApiKey middleware.
//
// Parameters:
// - router: The Gin router instance to configure.
//
// Example:
//   router := gin.Default()
//   routes.SetupAboutRoutes(router)
func SetupAboutRoutes(router *gin.Engine) {
	router.POST("/createAbout", middlewares.ValidateApiKey(), aboutcontrollers.CreateAbout)
	router.GET("/about", middlewares.ValidateApiKey(), aboutcontrollers.GetAbout)
}
