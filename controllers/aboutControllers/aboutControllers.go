// Package aboutcontrollers implements the API endpoints for managing "About" content.
package aboutcontrollers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"
    "strconv"

    "github.com/EkoAgustina/go-ms-portfolio/config/database"
    "github.com/EkoAgustina/go-ms-portfolio/models/aboutModels"
    "github.com/EkoAgustina/go-ms-portfolio/utils"

    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "gorm.io/gorm"
)

// ctx is the context used for Redis operations.
var ctx = context.Background()

// CreateAbout handles the HTTP request to create a new "About" entry.
// It expects a JSON body containing the About model data.
// On success, it responds with a 201 Created status and the created entry data.
// On failure (e.g., invalid JSON), it responds with a 400 Bad Request status.
func CreateAbout(c *gin.Context) {
    var about aboutmodels.About
    if err := c.ShouldBindJSON(&about); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "responseCode":    http.StatusBadRequest,
            "responseMessage": "Invalid request body format",
        })
        return
    }

    database.DB.Session(&gorm.Session{PrepareStmt: true}).Create(&about)

    c.JSON(http.StatusCreated, gin.H{
        "responseCode": http.StatusCreated,
        "data": about,
    })
}

// GetAbout handles the HTTP request to retrieve "About" entries.
// It accepts an optional query parameter "id" to fetch a specific entry.
// It uses Redis for caching; if data is not in the cache, it retrieves it from the database.
// On success, it responds with a 200 OK status and the requested data.
// If the entry is not found, it responds with a 404 Not Found status.
// In case of cache errors, it responds with a 500 Internal Server Error status.
func GetAbout(c *gin.Context) {
    var about []aboutmodels.About
    redisTtlStr := utils.LoadEnv("REDIS_CACHE_TTL")
    redisTtl, err := strconv.Atoi(redisTtlStr)
    if err != nil {
        log.Printf("Error converting REDIS_CACHE_TTL to integer: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "responseCode":    http.StatusInternalServerError,
            "responseMessage": "Invalid Redis TTL configuration",
        })
        return
    }

    id := c.Query("id")
    var cacheKey string
    if id != "" {
        cacheKey = "about:" + id
    } else {
        cacheKey = "about:all"
    }

    rdb, _ := c.Get("redis")
    redisClient := rdb.(*redis.Client)

    cachedData, err := redisClient.Get(ctx, cacheKey).Result()
    if err != nil {
        if err == redis.Nil {
            log.Printf("Cache miss for key %s", cacheKey)

            // Jika ada parameter id, ambil satu project, jika tidak, ambil semua project
            if id != "" {
                if err := database.DB.Session(&gorm.Session{PrepareStmt: true}).First(&about, id).Error; err != nil {
                    log.Printf("Error fetching from database: %v", err)
                    c.JSON(http.StatusNotFound, gin.H{
                        "responseCode":    http.StatusNotFound,
                        "responseMessage": "Content not found",
                    })
                    return
                }
            } else {
                // Ambil semua project jika tidak ada id
                if err := database.DB.Session(&gorm.Session{PrepareStmt: true}).Find(&about).Error; err != nil {
                    log.Printf("Error fetching from database: %v", err)
                    c.JSON(http.StatusInternalServerError, gin.H{
                        "responseCode":    http.StatusInternalServerError,
                        "responseMessage": "Error retrieving data",
                    })
                    return
                }
            }
            
            if len(about) == 0 {
                c.JSON(http.StatusNotFound, gin.H{
                    "responseCode":    http.StatusNotFound,
                    "responseMessage": "No content found",
                })
                return
            }

            response := gin.H{
                "responseCode": http.StatusOK,
                "data":         about,
            }

            jsonData, err := json.Marshal(response)
            if err != nil {
                log.Printf("Error marshaling JSON: %v", err)
                c.JSON(http.StatusInternalServerError, gin.H{
                    "responseCode":    http.StatusInternalServerError,
                    "responseMessage": "Internal Server Error",
                })
                return
            }

            err = redisClient.Set(ctx, cacheKey, jsonData, time.Duration(redisTtl)*time.Second).Err()
            if err != nil {
                log.Printf("Error saving to Redis: %v", err)
            } else {
                log.Printf("Key %s saved to Redis", cacheKey)
            }

            c.JSON(http.StatusOK, response)
        } else {
            log.Printf("Error accessing Redis: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "responseCode":    http.StatusInternalServerError,
                "responseMessage": "Error accessing cache",
            })
            return
        }
    } else {
        log.Printf("Cache hit for key %s", cacheKey)
        cachedResponse := gin.H{}
        if err := json.Unmarshal([]byte(cachedData), &cachedResponse); err != nil {
            log.Printf("Error unmarshaling JSON from Redis: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "responseCode":    http.StatusInternalServerError,
                "responseMessage": "Error processing cached data",
            })
            return
        }
        c.JSON(http.StatusOK, cachedResponse)
    }
}
