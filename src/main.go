package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sachinmangla/rentiflat/config"
	"github.com/sachinmangla/rentiflat/database"
	"github.com/sachinmangla/rentiflat/server"
)

func main() {
	fmt.Println("Starting Server...")

	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	// TODO: Add a retry mechanism
	for i := 0; i < 5; i++ {
		err := database.DatabaseCon(database.NewDatabaseConfig())
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	// check if the database is ready
	if err := database.CheckDBConnection(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform automatic database migrations
	database.MigrateDB()

	// Start the server
	port := config.GetEnv("APP_PORT", "8080")
	if err := server.RunServer(port); err != nil {
		log.Fatalf("Failed to start Server on port %s: %v", port, err)
	}

	fmt.Printf("Server started on port %s\n", port)
}
