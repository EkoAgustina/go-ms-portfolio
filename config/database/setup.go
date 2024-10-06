package database

import (
	"fmt"
	"github.com/EkoAgustina/go-ms-portfolio/models/projectModels"
	"github.com/EkoAgustina/go-ms-portfolio/models/aboutModels"
	"github.com/EkoAgustina/go-ms-portfolio/models/contactModels"
	"github.com/EkoAgustina/go-ms-portfolio/utils"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection instance.
var DB *gorm.DB

// Connect connects to the PostgreSQL database and performs automatic migrations.
// It reads configuration from environment variables and retries the connection up to 5 times if it fails.
// If successful, it logs a confirmation message. If it fails after retries, it logs a fatal error and exits.
//
// Environment Variables:
// - DB_HOST: Database server address.
// - DB_USER: Database username.
// - DB_PASSWORD: Database password.
// - DB_NAME: Name of the database.
// - DB_PORT: Port number for the database server.
//
// Example:
//   database.Connect()
func Connect () {
	host := utils.LoadEnv("DB_HOST")
	user := utils.LoadEnv("DB_USER")
	password := utils.LoadEnv("DB_PASSWORD")
	dbname := utils.LoadEnv("DB_NAME")
	port := utils.LoadEnv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Database connection established")
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	DB.AutoMigrate(&aboutmodels.About{}, &projectmodels.Project{}, &contactmodels.Contact{})
}