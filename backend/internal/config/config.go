package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Environment string
	Port        string

	// Supabase Configuration
	SupabaseURL        string
	SupabaseAnonKey    string
	SupabaseServiceKey string

	// JWT Configuration
	JWTSecret string
	JWTExpiry string

	// Payment Gateway Configuration
	RazorpayKeyID     string
	RazorpayKeySecret string
	PaystackSecretKey string
	PaystackPublicKey string

	// Coin System Configuration
	WelcomeCoins        int
	MinCoinsForPurchase int

	// Audio Storage Configuration
	AudioBucketName  string
	MaxAudioFileSize string

	// CORS Configuration
	AllowedOrigins []string
}

func Load() *Config {
	return &Config{
		Environment:         getEnv("ENV", "development"),
		Port:                getEnv("PORT", "3003"),
		SupabaseURL:         getEnv("SUPABASE_URL", ""),
		SupabaseAnonKey:     getEnv("SUPABASE_ANON_KEY", ""),
		SupabaseServiceKey:  getEnv("SUPABASE_SERVICE_ROLE_KEY", ""),
		JWTSecret:           getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiry:           getEnv("JWT_EXPIRY", "24h"),
		RazorpayKeyID:       getEnv("RAZORPAY_KEY_ID", ""),
		RazorpayKeySecret:   getEnv("RAZORPAY_KEY_SECRET", ""),
		PaystackSecretKey:   getEnv("PAYSTACK_SECRET_KEY", ""),
		PaystackPublicKey:   getEnv("PAYSTACK_PUBLIC_KEY", ""),
		WelcomeCoins:        getEnvAsInt("WELCOME_COINS", 50),
		MinCoinsForPurchase: getEnvAsInt("MIN_COINS_FOR_PURCHASE", 10),
		AudioBucketName:     getEnv("AUDIO_BUCKET_NAME", "audio-episodes"),
		MaxAudioFileSize:    getEnv("MAX_AUDIO_FILE_SIZE", "100MB"),
		AllowedOrigins:      getEnvAsSlice("ALLOWED_ORIGINS", []string{"http://localhost:3004", "http://localhost:3003"}),
	}
}

// IsDirectDatabaseURL checks if the Supabase URL is a direct database connection
func (c *Config) IsDirectDatabaseURL() bool {
	return strings.HasPrefix(c.SupabaseURL, "postgresql://")
}

// GetDatabaseURL returns the database connection URL
func (c *Config) GetDatabaseURL() string {
	// First, check if DATABASE_URL is set (direct connection)
	if dbURL := getEnv("DATABASE_URL", ""); dbURL != "" {
		return dbURL
	}

	if c.IsDirectDatabaseURL() {
		return c.SupabaseURL
	}

	// Convert REST API URL to direct database URL
	// Extract project reference from https://your-project.supabase.co
	if strings.HasPrefix(c.SupabaseURL, "https://") {
		parts := strings.Split(c.SupabaseURL, "//")
		if len(parts) == 2 {
			projectRef := strings.Split(parts[1], ".")[0]
			// You'll need to get the database password from environment or config
			dbPassword := getEnv("SUPABASE_DB_PASSWORD", "")
			if dbPassword != "" {
				return fmt.Sprintf("postgresql://postgres:%s@db.%s.supabase.co:5432/postgres", dbPassword, projectRef)
			}
		}
	}

	return ""
}

// GetSupabaseProjectRef extracts the project reference from the URL
func (c *Config) GetSupabaseProjectRef() string {
	if strings.HasPrefix(c.SupabaseURL, "https://") {
		parts := strings.Split(c.SupabaseURL, "//")
		if len(parts) == 2 {
			return strings.Split(parts[1], ".")[0]
		}
	}
	return ""
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
