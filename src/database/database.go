package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sachinmangla/rentiflat/config"
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
	port, _ := strconv.Atoi(config.GetEnv("DB_PORT", "5432"))

	fmt.Println("DB_HOST", config.GetEnv("DB_HOST", "localhost"))

	return &DatabaseConfig{
		Host:     config.GetEnv("DB_HOST", "localhost"),
		Port:     port,
		User:     config.GetEnv("DB_USER", "postgres"),
		Password: config.GetEnv("DB_PASSWORD", ""),
		DBName:   config.GetEnv("DB_NAME", "postgres"),
		SSLMode:  config.GetEnv("DB_SSLMODE", "disable"),
	}
}

var Db *gorm.DB

// DatabaseCon establishes connection to the database
func DatabaseCon(config *DatabaseConfig) error {
	// Connect to the database
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.DBName, config.Password, config.SSLMode)
	db, err := gorm.Open("postgres", connectionString)
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
