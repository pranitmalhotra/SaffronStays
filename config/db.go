package config

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"github.com/joho/godotenv"
)

// ConnectDB establishes a connection to the PostgreSQL database using GORM.
// It loads environment variables from a .env file and uses the DB_URL environment
// variable to connect to the database. It returns a pointer to the GORM DB instance
// and an error if the connection fails.
func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbURL,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database using GORM!")
	return db, nil
}