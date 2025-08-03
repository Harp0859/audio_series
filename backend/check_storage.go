package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Supabase configuration
	supabaseURL := "https://mhbcihpkcetbzdrzciqe.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im1oYmNpaHBrY2V0YnpkcnpjaXFlIiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTc1MzM3NjQ1NSwiZXhwIjoyMDY4OTUyNDU1fQ.PGGhJx3wlTBfSNhaQf1XOw6In21vCwgDfcdTnIsIDdY"

	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Println("🔍 Checking your Supabase storage...")

	// First, let's check what buckets exist
	fmt.Println("\n📦 Checking available buckets...")

	req, _ := http.NewRequest("GET", supabaseURL+"/storage/v1/bucket/list", nil)
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Failed to list buckets: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("📊 Buckets Response Status: %s\n", resp.Status)

	if resp.StatusCode == 200 {
		var buckets []map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&buckets); err != nil {
			fmt.Printf("❌ Failed to decode buckets response: %v\n", err)
			return
		}

		fmt.Printf("✅ Found %d buckets:\n", len(buckets))
		for i, bucket := range buckets {
			name := bucket["name"]
			public := bucket["public"]
			fmt.Printf("   %d. %s (Public: %v)\n", i+1, name, public)
		}

		// Now check files in the audio bucket
		fmt.Println("\n🎵 Checking files in 'audio' bucket...")

		req2, _ := http.NewRequest("GET", supabaseURL+"/storage/v1/object/list/audio", nil)
		req2.Header.Set("apikey", supabaseKey)
		req2.Header.Set("Authorization", "Bearer "+supabaseKey)

		resp2, err := client.Do(req2)
		if err != nil {
			fmt.Printf("❌ Failed to list files: %v\n", err)
			return
		}
		defer resp2.Body.Close()

		fmt.Printf("📊 Files Response Status: %s\n", resp2.Status)

		if resp2.StatusCode == 200 {
			var result map[string]interface{}
			if err := json.NewDecoder(resp2.Body).Decode(&result); err != nil {
				fmt.Printf("❌ Failed to decode files response: %v\n", err)
				return
			}

			if files, ok := result["data"].([]interface{}); ok {
				fmt.Printf("✅ Found %d files in audio bucket:\n", len(files))
				for i, file := range files {
					if fileMap, ok := file.(map[string]interface{}); ok {
						name := fileMap["name"]
						size := fileMap["metadata"].(map[string]interface{})["size"]
						updatedAt := fileMap["updated_at"]
						fmt.Printf("   %d. %s (Size: %v bytes, Updated: %s)\n", i+1, name, size, updatedAt)
					}
				}

				// Generate public URLs for each file
				if len(files) > 0 {
					fmt.Println("\n🔗 Public URLs for your files:")
					for i, file := range files {
						if fileMap, ok := file.(map[string]interface{}); ok {
							name := fileMap["name"]
							publicURL := fmt.Sprintf("%s/storage/v1/object/public/audio/%s", supabaseURL, name)
							fmt.Printf("   %d. %s\n      URL: %s\n", i+1, name, publicURL)
						}
					}
				}
			} else {
				fmt.Println("⚠️  No files found or unexpected response format")
			}
		} else {
			fmt.Printf("❌ Failed to list files: %s\n", resp2.Status)
			fmt.Println("\n💡 The bucket might be private or there might be an issue with the API call.")
		}
	} else {
		fmt.Printf("❌ Failed to list buckets: %s\n", resp.Status)
	}

	fmt.Println("\n🎉 Storage check completed!")
}
