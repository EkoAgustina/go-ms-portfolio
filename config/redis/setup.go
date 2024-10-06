package redis

import (
	"context"
	"fmt"

	"github.com/EkoAgustina/go-ms-portfolio/utils"
	"github.com/go-redis/redis/v8"
)


// RedisConfig holds the configuration details for connecting to a Redis server.
type RedisConfig struct {
	Address  string // Redis server address
	Port     string // Redis server port
	Password string // Redis server password (if any)
	DB       int    // Redis database number
}

// LoadRedisConfig loads Redis configuration from environment variables.
// It returns a RedisConfig object with the configuration values.
// Returns an error if the necessary environment variables are not set.
func LoadRedisConfig() (*RedisConfig, error) {
	address := utils.LoadEnv("REDIS_HOST")
	port := utils.LoadEnv("REDIS_PORT")
	password := utils.LoadEnv("REDIS_PASSWORD")

	
	if address == "" || port == "" {
		return nil, fmt.Errorf("REDIS_HOST and REDIS_PORT must be set")
	}

	return &RedisConfig{
		Address:  address,
		Port:     port,
		Password: password,
		DB:       1, // Default Redis database number
	}, nil
}

// SetupRedis initializes and returns a Redis client based on the loaded configuration.
// It takes a context as a parameter for connection management.
// Returns a redis.Client object and an error, if any occurs during connection setup.
func SetupRedis(ctx context.Context) (*redis.Client, error) {
	config, err := LoadRedisConfig()
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
		
	}

	return rdb, nil
}
