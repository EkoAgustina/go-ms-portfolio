// Package middlewares contains custom middleware functions for the Gin web framework.
package middlewares

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/EkoAgustina/go-ms-portfolio/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// ValidateApiKey checks for the presence and validity of an API key in the request header.
// If the API key is missing or invalid, it responds with a 403 Forbidden status and aborts the request.
// If the API key is valid, it allows the request to proceed to the next handler.
//
// Returns a gin.HandlerFunc that can be used as middleware.
func ValidateApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("x-api-key")

		if apikey == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"responseCode": http.StatusForbidden,
				"error": "apikey required",
			})
			c.Abort()
			return
		}

		if apikey != utils.LoadEnv("API_KEY") {
			c.JSON(http.StatusForbidden, gin.H{
				"responseCode": http.StatusForbidden,
				"error": "Invalid apikey",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CustomWriter is a custom ResponseWriter that captures the response body for logging.
type CustomWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write writes the response body and captures it in the CustomWriter.
func (w CustomWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Save the response body
	return w.ResponseWriter.Write(b) // Send the response to the client
}

// CustomLogger logs request and response details, including headers, body, and execution time.
// It wraps the request processing to capture data for logging purposes.
//
// Returns a gin.HandlerFunc that can be used as middleware.
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request headers
		log.Printf("Request Headers: %v", c.Request.Header)

		// Log request body
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		log.Printf("Request Body: %s", string(bodyBytes))

		// Start timer to record the request duration
		startTime := time.Now()

		// Capture response using custom writer
		customWriter := &CustomWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = customWriter

		// Proceed to the next handler
		c.Next()

		// Log response headers
		log.Printf("Response Headers: %v", c.Writer.Header())

		// Log response body
		log.Printf("Response Body: %s", customWriter.body.String())

		// Log execution time
		duration := time.Since(startTime)
		log.Printf("Request processed in %s", duration)
	}
}

// RedisMiddleware sets up the Redis client in the context for later use in handlers.
// It allows handlers to access the Redis client without needing to pass it explicitly.
//
// Parameters:
// - rdb: The Redis client instance to set in the context.
//
// Returns a gin.HandlerFunc that can be used as middleware.
func RedisMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", rdb)
		c.Next()
	}
}
