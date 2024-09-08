package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DatabaseConfig represents database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDatabaseConfig initializes and returns a new DatabaseConfig object
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "Sac@2121",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
}

var Db *gorm.DB

// DatabaseCon establishes connection to the database
func DatabaseCon(config *DatabaseConfig) error {
	// Connect to the database
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.DBName, config.Password, config.SSLMode))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Ping the database to check if it's responsive
	if err := db.DB().Ping(); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}
	Db = db
	log.Println("Database connection established successfully")
	return nil
}

// MigrateDB performs automatic migrations for database schema
func MigrateDB() {
	Db.AutoMigrate(&OwnerDetails{})
	Db.AutoMigrate(&FlatDetails{})
}

func GetDb() *gorm.DB {
	return Db
}
