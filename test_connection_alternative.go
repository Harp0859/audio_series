package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get database URL
	dbURL := os.Getenv("SUPABASE_URL")
	if dbURL == "" {
		log.Fatal("SUPABASE_URL is required")
	}

	fmt.Println("🔗 Testing Supabase connection with alternative parameters...")
	fmt.Printf("URL: %s\n", dbURL)

	// Try different connection parameters
	connectionAttempts := []string{
		dbURL,
		dbURL + "?sslmode=require",
		dbURL + "?sslmode=prefer",
		dbURL + "?sslmode=disable",
	}

	for i, connStr := range connectionAttempts {
		fmt.Printf("\n🔄 Attempt %d: %s\n", i+1, connStr)

		// Connect to database
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("❌ Failed to open connection: %v\n", err)
			continue
		}

		// Set connection timeout
		db.SetConnMaxLifetime(10 * time.Second)
		db.SetMaxOpenConns(1)

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			fmt.Printf("❌ Failed to ping database: %v\n", err)
			db.Close()
			continue
		}

		fmt.Printf("✅ Successfully connected with attempt %d!\n", i+1)
		db.Close()
		return
	}

	fmt.Println("\n❌ All connection attempts failed")
	fmt.Println("\n💡 Troubleshooting tips:")
	fmt.Println("1. Check your internet connection")
	fmt.Println("2. Verify your Supabase project is active")
	fmt.Println("3. Try using a VPN if you're behind a corporate firewall")
	fmt.Println("4. Check if port 5432 is blocked by your network")
}
