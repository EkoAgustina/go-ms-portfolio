// Package routes contains the configuration for application routes.
package routes

import (
	"github.com/EkoAgustina/go-ms-portfolio/controllers/contactControllers"
	"github.com/EkoAgustina/go-ms-portfolio/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupContactRoutes configures routes for "contact" endpoints on the Gin router.
// This function sets up the following routes:
// - POST /contactme: Creates a new contact entry. Validated with ValidateApiKey middleware.
// - GET /contactme: Retrieves contact entries. Validated with ValidateApiKey middleware.
//
// Parameters:
// - router: The Gin router instance to configure.
//
// Example:
//   router := gin.Default()
//   routes.SetupContactRoutes(router)
func SetupContactRoutes(router *gin.Engine) {
	router.POST("/contactme", middlewares.ValidateApiKey(), contactcontrollers.CreateContact)
	router.GET("/contactme", middlewares.ValidateApiKey(), contactcontrollers.GetContactMe)
}
