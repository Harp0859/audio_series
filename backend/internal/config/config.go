package config

import (
	"os"
	"strconv"
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
		// Simple comma-separated values parsing
		// In production, you might want more sophisticated parsing
		return []string{value}
	}
	return defaultValue
}
