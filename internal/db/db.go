package db

import (
	"fmt"
	"log"
	"url-shortner/internal/config"
	"url-shortner/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db *gorm.DB
}

var (
	database   = config.DB_NAME
	password   = config.DB_PASSWORD
	username   = config.DB_USER
	port       = config.DB_PORT
	host       = config.DB_HOST
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.URL{})

	log.Println("Connected to database")

	dbInstance = &Service{
		DB: db,
	}
	return dbInstance
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	dbInstance, err := s.db.DB()
	if err != nil {
		return err
	}
	dbInstance.Close()
	return nil
}
