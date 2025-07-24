package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	Phone       *string   `json:"phone,omitempty" db:"phone"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	AvatarURL   *string   `json:"avatar_url,omitempty" db:"avatar_url"`
	CoinBalance int       `json:"coin_balance" db:"coin_balance"`
	Role        string    `json:"role" db:"role"` // user, admin
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Series represents an audio series
type Series struct {
	ID            uuid.UUID `json:"id" db:"id"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	CoverImage    string    `json:"cover_image" db:"cover_image"`
	Author        string    `json:"author" db:"author"`
	Category      string    `json:"category" db:"category"`
	IsPremium     bool      `json:"is_premium" db:"is_premium"`
	TotalEpisodes int       `json:"total_episodes" db:"total_episodes"`
	CreatedBy     uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Episode represents an individual episode in a series
type Episode struct {
	ID            uuid.UUID `json:"id" db:"id"`
	SeriesID      uuid.UUID `json:"series_id" db:"series_id"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	AudioURL      string    `json:"audio_url" db:"audio_url"`
	Duration      int       `json:"duration" db:"duration"` // in seconds
	EpisodeNumber int       `json:"episode_number" db:"episode_number"`
	CoinPrice     int       `json:"coin_price" db:"coin_price"`
	IsLocked      bool      `json:"is_locked" db:"is_locked"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Purchase represents a user's purchase of an episode or series
type Purchase struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	EpisodeID *uuid.UUID `json:"episode_id,omitempty" db:"episode_id"`
	SeriesID  *uuid.UUID `json:"series_id,omitempty" db:"series_id"`
	Type      string     `json:"type" db:"type"`     // episode, series, coins
	Amount    int        `json:"amount" db:"amount"` // coins spent
	PaymentID *string    `json:"payment_id,omitempty" db:"payment_id"`
	Status    string     `json:"status" db:"status"` // completed, pending, failed
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
}

// CoinTransaction represents a coin balance change
type CoinTransaction struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Type        string    `json:"type" db:"type"`       // purchase, welcome, refund, admin
	Amount      int       `json:"amount" db:"amount"`   // positive for credit, negative for debit
	Balance     int       `json:"balance" db:"balance"` // balance after transaction
	Description string    `json:"description" db:"description"`
	ReferenceID *string   `json:"reference_id,omitempty" db:"reference_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Payment represents a payment transaction
type Payment struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Amount      int       `json:"amount" db:"amount"` // in smallest currency unit
	Currency    string    `json:"currency" db:"currency"`
	Coins       int       `json:"coins" db:"coins"`
	Gateway     string    `json:"gateway" db:"gateway"` // razorpay, paystack
	GatewayRef  string    `json:"gateway_ref" db:"gateway_ref"`
	Status      string    `json:"status" db:"status"`             // pending, completed, failed
	PaymentData string    `json:"payment_data" db:"payment_data"` // JSON string
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// AuthRequest represents authentication request
type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone,omitempty"`
}

// SeriesWithEpisodes represents a series with its episodes
type SeriesWithEpisodes struct {
	Series   *Series    `json:"series"`
	Episodes []*Episode `json:"episodes"`
}

// EpisodeWithPurchase represents an episode with purchase status
type EpisodeWithPurchase struct {
	Episode   *Episode `json:"episode"`
	IsOwned   bool     `json:"is_owned"`
	CanUnlock bool     `json:"can_unlock"`
}

// CoinBundle represents available coin bundles for purchase
type CoinBundle struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Coins     int       `json:"coins" db:"coins"`
	Price     int       `json:"price" db:"price"` // in smallest currency unit
	Currency  string    `json:"currency" db:"currency"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// PaymentRequest represents a payment initiation request
type PaymentRequest struct {
	BundleID string `json:"bundle_id" binding:"required"`
	Currency string `json:"currency" binding:"required"` // INR, NGN
}

// PaymentResponse represents payment gateway response
type PaymentResponse struct {
	PaymentID   string `json:"payment_id"`
	GatewayRef  string `json:"gateway_ref"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Gateway     string `json:"gateway"`
	RedirectURL string `json:"redirect_url,omitempty"`
}

// AdminStats represents admin dashboard statistics
type AdminStats struct {
	TotalUsers     int `json:"total_users"`
	TotalSeries    int `json:"total_series"`
	TotalEpisodes  int `json:"total_episodes"`
	TotalRevenue   int `json:"total_revenue"`
	MonthlyRevenue int `json:"monthly_revenue"`
	ActiveUsers    int `json:"active_users"`
}
