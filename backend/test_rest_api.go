package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	// Get Supabase URL and key
	supabaseURL := os.Getenv("SUPABASE_URL")
	anonKey := os.Getenv("SUPABASE_ANON_KEY")

	if supabaseURL == "" || anonKey == "" {
		fmt.Println("❌ SUPABASE_URL and SUPABASE_ANON_KEY are required")
		return
	}

	fmt.Println("🔗 Testing Supabase REST API connection...")
	fmt.Printf("URL: %s\n", supabaseURL)
	fmt.Printf("Key: %s...\n", anonKey[:20])

	// Test REST API connection
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Try to access the health endpoint
	req, err := http.NewRequest("GET", supabaseURL+"/rest/v1/", nil)
	if err != nil {
		fmt.Printf("❌ Failed to create request: %v\n", err)
		return
	}

	req.Header.Set("apikey", anonKey)
	req.Header.Set("Authorization", "Bearer "+anonKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Failed to connect to REST API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Failed to read response: %v\n", err)
		return
	}

	fmt.Printf("✅ REST API Response Status: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n", string(body))

	if resp.StatusCode == 200 {
		fmt.Println("🎉 REST API connection successful!")
		fmt.Println("\n💡 Since direct database connection is failing, you can:")
		fmt.Println("1. Use Supabase REST API for development")
		fmt.Println("2. Check your network/firewall settings")
		fmt.Println("3. Contact your network administrator about port 5432")
		fmt.Println("4. Try using a different network (mobile hotspot, etc.)")
	} else {
		fmt.Printf("⚠️  REST API returned status %d\n", resp.StatusCode)
	}
}
