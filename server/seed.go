//go:build seed

package main

import (
	"log"
	"riddles-server/database"
	"riddles-server/utils"
)

func main() {
	// Initialize database
	database.ConnectDB()
	database.MigrateDB()

	// Seed database
	utils.SeedDatabase()

	log.Println("Database seeding completed successfully")
}