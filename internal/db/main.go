package database

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// ent "github.com/thisjustinf/togo-list/internal/db/entities"
	"github.com/thisjustinf/togo-list/internal/graphql/models"
)

// var DBInstance *gorm.DB
// var err error

func InitDB() (*gorm.DB, error) {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	log.Println("Successfully connected to the DB!")
	db.AutoMigrate(&models.User{}, &models.Todo{})
	return db, nil
}
