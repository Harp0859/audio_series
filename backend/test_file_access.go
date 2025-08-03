package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Supabase configuration
	supabaseURL := "https://mhbcihpkcetbzdrzciqe.supabase.co"

	client := &http.Client{Timeout: 10 * time.Second}

	fmt.Println("🔍 Testing direct file access...")

	// Test some common file extensions
	testFiles := []string{
		"sample.mp4",
		"video.mp4",
		"test.mp4",
		"audio.mp3",
		"sample.mp3",
		"test.mp3",
		"file.mp4",
		"myvideo.mp4",
	}

	fmt.Println("\n🔗 Testing public URLs for common filenames...")

	for _, filename := range testFiles {
		url := fmt.Sprintf("%s/storage/v1/object/public/audio/%s", supabaseURL, filename)

		resp, err := client.Head(url)
		if err != nil {
			fmt.Printf("❌ %s - Error: %v\n", filename, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("✅ %s - FOUND! (Status: %s)\n", filename, resp.Status)
			fmt.Printf("   URL: %s\n", url)
		} else if resp.StatusCode == 404 {
			fmt.Printf("❌ %s - Not found (404)\n", filename)
		} else {
			fmt.Printf("⚠️  %s - Status: %s\n", filename, resp.Status)
		}
	}

	fmt.Println("\n💡 If none of these work, please:")
	fmt.Println("1. Go to your Supabase dashboard")
	fmt.Println("2. Navigate to Storage → audio bucket")
	fmt.Println("3. Tell me the exact filename of your MP4 file")
	fmt.Println("4. I'll update the frontend with the correct URL")
}
