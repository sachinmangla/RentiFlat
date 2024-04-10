package main

import (
	"fmt"
	"log"

	"github.com/sachinmangla/rentiflat/database"
	"github.com/sachinmangla/rentiflat/server"
)

func main() {
	fmt.Println("Starting Server...")

	// Connect to the database
	db, err := database.DatabaseCon(database.NewDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform automatic database migrations
	database.MigrateDB(db)

	// Start the server
	port := "8080"
	if err := server.RunServer(port); err != nil {
		log.Fatalf("Failed to start Server on port %s: %v", port, err)
	}

	fmt.Printf("Server started on port %s\n", port)
}
