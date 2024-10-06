package main

import (
	"context"
	"log"
	"net/http"

	"github.com/EkoAgustina/go-ms-portfolio/config/database"
	"github.com/EkoAgustina/go-ms-portfolio/config/redis"
	"github.com/EkoAgustina/go-ms-portfolio/routes"
	"github.com/EkoAgustina/go-ms-portfolio/middlewares"
	"github.com/EkoAgustina/go-ms-portfolio/utils"
	"github.com/gin-gonic/gin"
)
var ctx = context.Background()
func main () {
	database.Connect()
	// Set up Redis
	rdb, err := redis.SetupRedis(ctx)
	if err != nil {
		log.Fatalf("Failed to set up Redis: %v", err)
	}

	// Check Redis connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error while pinging Redis: %v", err)
	}

	log.Println("Successfully connected to Redis", pong)

	router := gin.Default()
	router.Use(middlewares.CustomLogger())
	router.Use(middlewares.RedisMiddleware(rdb))

	routes.SetupAboutRoutes(router)
	routes.SetupProjectRoutes(router)
	routes.SetupContactRoutes(router)

	log.Println(http.ListenAndServe(":"+utils.LoadEnv("GO_PORT"), router))
}