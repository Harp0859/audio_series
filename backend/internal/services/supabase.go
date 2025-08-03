package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"audio-series-app/backend/internal/config"
	"audio-series-app/backend/internal/models"

	"github.com/google/uuid"
)

type SupabaseService struct {
	client  *http.Client
	config  *config.Config
	baseURL string
	apiKey  string
}

func NewSupabaseService(cfg *config.Config) *SupabaseService {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	log.Println("✅ Successfully initialized Supabase service")

	return &SupabaseService{
		client:  client,
		config:  cfg,
		baseURL: cfg.SupabaseURL + "/rest/v1",
		apiKey:  cfg.SupabaseServiceKey,
	}
}

// makeRequest makes an HTTP request to Supabase REST API
func (s *SupabaseService) makeRequest(ctx context.Context, method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, s.baseURL+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("apikey", s.apiKey)
	req.Header.Set("Authorization", "Bearer "+s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// User operations
func (s *SupabaseService) CreateUser(ctx context.Context, user *models.User) error {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "user"
	user.IsActive = true
	user.CoinBalance = s.config.WelcomeCoins

	_, err := s.makeRequest(ctx, "POST", "/users", user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Create welcome coin transaction
	transaction := &models.CoinTransaction{
		UserID:      user.ID,
		Type:        "welcome",
		Amount:      s.config.WelcomeCoins,
		Balance:     s.config.WelcomeCoins,
		Description: "Welcome bonus coins",
	}

	return s.CreateCoinTransaction(ctx, transaction)
}

func (s *SupabaseService) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, email, phone, first_name, last_name, avatar_url, coin_balance, role, is_active, created_at, updated_at
		FROM users WHERE id = $1
	`

	user := &models.User{}
	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID, &user.Email, &user.Phone, &user.FirstName, &user.LastName,
		&user.AvatarURL, &user.CoinBalance, &user.Role, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return user, nil
}

func (s *SupabaseService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, email, phone, first_name, last_name, avatar_url, coin_balance, role, is_active, created_at, updated_at
		FROM users WHERE email = $1
	`

	user := &models.User{}
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Phone, &user.FirstName, &user.LastName,
		&user.AvatarURL, &user.CoinBalance, &user.Role, &user.IsActive,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	return user, nil
}

func (s *SupabaseService) UpdateUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET email = $2, phone = $3, first_name = $4, last_name = $5, 
		    avatar_url = $6, coin_balance = $7, role = $8, is_active = $9, updated_at = $10
		WHERE id = $1
	`

	user.UpdatedAt = time.Now()
	_, err := s.db.ExecContext(ctx, query,
		user.ID, user.Email, user.Phone, user.FirstName, user.LastName,
		user.AvatarURL, user.CoinBalance, user.Role, user.IsActive, user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

// Series operations
func (s *SupabaseService) CreateSeries(ctx context.Context, series *models.Series) error {
	query := `
		INSERT INTO series (id, title, description, cover_image, author, category, is_premium, total_episodes, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	series.ID = uuid.New()
	series.CreatedAt = time.Now()
	series.UpdatedAt = time.Now()

	_, err := s.db.ExecContext(ctx, query,
		series.ID, series.Title, series.Description, series.CoverImage,
		series.Author, series.Category, series.IsPremium, series.TotalEpisodes,
		series.CreatedBy, series.CreatedAt, series.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create series: %v", err)
	}

	return nil
}

func (s *SupabaseService) GetSeries(ctx context.Context) ([]*models.Series, error) {
	query := `
		SELECT id, title, description, cover_image, author, category, is_premium, total_episodes, created_by, created_at, updated_at
		FROM series ORDER BY created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get series: %v", err)
	}
	defer rows.Close()

	var series []*models.Series
	for rows.Next() {
		s := &models.Series{}
		err := rows.Scan(
			&s.ID, &s.Title, &s.Description, &s.CoverImage, &s.Author,
			&s.Category, &s.IsPremium, &s.TotalEpisodes, &s.CreatedBy,
			&s.CreatedAt, &s.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan series: %v", err)
		}
		series = append(series, s)
	}

	return series, nil
}

func (s *SupabaseService) GetSeriesByID(ctx context.Context, seriesID uuid.UUID) (*models.Series, error) {
	query := `
		SELECT id, title, description, cover_image, author, category, is_premium, total_episodes, created_by, created_at, updated_at
		FROM series WHERE id = $1
	`

	series := &models.Series{}
	err := s.db.QueryRowContext(ctx, query, seriesID).Scan(
		&series.ID, &series.Title, &series.Description, &series.CoverImage,
		&series.Author, &series.Category, &series.IsPremium, &series.TotalEpisodes,
		&series.CreatedBy, &series.CreatedAt, &series.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get series: %v", err)
	}

	return series, nil
}

// Episode operations
func (s *SupabaseService) CreateEpisode(ctx context.Context, episode *models.Episode) error {
	query := `
		INSERT INTO episodes (id, series_id, title, description, audio_url, duration, episode_number, coin_price, is_locked, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	episode.ID = uuid.New()
	episode.CreatedAt = time.Now()
	episode.UpdatedAt = time.Now()

	_, err := s.db.ExecContext(ctx, query,
		episode.ID, episode.SeriesID, episode.Title, episode.Description,
		episode.AudioURL, episode.Duration, episode.EpisodeNumber,
		episode.CoinPrice, episode.IsLocked, episode.CreatedAt, episode.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create episode: %v", err)
	}

	return nil
}

func (s *SupabaseService) GetEpisodesBySeriesID(ctx context.Context, seriesID uuid.UUID) ([]*models.Episode, error) {
	query := `
		SELECT id, series_id, title, description, audio_url, duration, episode_number, coin_price, is_locked, created_at, updated_at
		FROM episodes WHERE series_id = $1 ORDER BY episode_number
	`

	rows, err := s.db.QueryContext(ctx, query, seriesID)
	if err != nil {
		return nil, fmt.Errorf("failed to get episodes: %v", err)
	}
	defer rows.Close()

	var episodes []*models.Episode
	for rows.Next() {
		episode := &models.Episode{}
		err := rows.Scan(
			&episode.ID, &episode.SeriesID, &episode.Title, &episode.Description,
			&episode.AudioURL, &episode.Duration, &episode.EpisodeNumber,
			&episode.CoinPrice, &episode.IsLocked, &episode.CreatedAt, &episode.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan episode: %v", err)
		}
		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func (s *SupabaseService) GetEpisodeByID(ctx context.Context, episodeID uuid.UUID) (*models.Episode, error) {
	query := `
		SELECT id, series_id, title, description, audio_url, duration, episode_number, coin_price, is_locked, created_at, updated_at
		FROM episodes WHERE id = $1
	`

	episode := &models.Episode{}
	err := s.db.QueryRowContext(ctx, query, episodeID).Scan(
		&episode.ID, &episode.SeriesID, &episode.Title, &episode.Description,
		&episode.AudioURL, &episode.Duration, &episode.EpisodeNumber,
		&episode.CoinPrice, &episode.IsLocked, &episode.CreatedAt, &episode.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get episode: %v", err)
	}

	return episode, nil
}

// Purchase operations
func (s *SupabaseService) CreatePurchase(ctx context.Context, purchase *models.Purchase) error {
	query := `
		INSERT INTO purchases (id, user_id, episode_id, series_id, type, amount, payment_id, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	purchase.ID = uuid.New()
	purchase.CreatedAt = time.Now()
	purchase.Status = "completed"

	_, err := s.db.ExecContext(ctx, query,
		purchase.ID, purchase.UserID, purchase.EpisodeID, purchase.SeriesID,
		purchase.Type, purchase.Amount, purchase.PaymentID, purchase.Status, purchase.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create purchase: %v", err)
	}

	return nil
}

func (s *SupabaseService) GetUserPurchases(ctx context.Context, userID uuid.UUID) ([]*models.Purchase, error) {
	query := `
		SELECT id, user_id, episode_id, series_id, type, amount, payment_id, status, created_at
		FROM purchases WHERE user_id = $1 ORDER BY created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get purchases: %v", err)
	}
	defer rows.Close()

	var purchases []*models.Purchase
	for rows.Next() {
		purchase := &models.Purchase{}
		err := rows.Scan(
			&purchase.ID, &purchase.UserID, &purchase.EpisodeID, &purchase.SeriesID,
			&purchase.Type, &purchase.Amount, &purchase.PaymentID, &purchase.Status, &purchase.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan purchase: %v", err)
		}
		purchases = append(purchases, purchase)
	}

	return purchases, nil
}

func (s *SupabaseService) HasUserPurchasedEpisode(ctx context.Context, userID, episodeID uuid.UUID) (bool, error) {
	query := `
		SELECT COUNT(*) FROM purchases 
		WHERE user_id = $1 AND episode_id = $2 AND status = 'completed'
	`

	var count int
	err := s.db.QueryRowContext(ctx, query, userID, episodeID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check purchase: %v", err)
	}

	return count > 0, nil
}

// Coin operations
func (s *SupabaseService) UpdateUserCoins(ctx context.Context, userID uuid.UUID, amount int) error {
	query := `
		UPDATE users SET coin_balance = coin_balance + $2, updated_at = NOW()
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, userID, amount)
	if err != nil {
		return fmt.Errorf("failed to update user coins: %v", err)
	}

	return nil
}

func (s *SupabaseService) CreateCoinTransaction(ctx context.Context, transaction *models.CoinTransaction) error {
	query := `
		INSERT INTO coin_transactions (id, user_id, type, amount, balance, description, reference_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	transaction.ID = uuid.New()
	transaction.CreatedAt = time.Now()

	_, err := s.db.ExecContext(ctx, query,
		transaction.ID, transaction.UserID, transaction.Type, transaction.Amount,
		transaction.Balance, transaction.Description, transaction.ReferenceID, transaction.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create coin transaction: %v", err)
	}

	return nil
}

// Payment operations
func (s *SupabaseService) CreatePayment(ctx context.Context, payment *models.Payment) error {
	query := `
		INSERT INTO payments (id, user_id, amount, currency, coins, gateway, gateway_ref, status, payment_data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	payment.ID = uuid.New()
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	_, err := s.db.ExecContext(ctx, query,
		payment.ID, payment.UserID, payment.Amount, payment.Currency,
		payment.Coins, payment.Gateway, payment.GatewayRef, payment.Status,
		payment.PaymentData, payment.CreatedAt, payment.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create payment: %v", err)
	}

	return nil
}

func (s *SupabaseService) UpdatePayment(ctx context.Context, paymentID uuid.UUID, status string, paymentData string) error {
	query := `
		UPDATE payments SET status = $2, payment_data = $3, updated_at = NOW()
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, paymentID, status, paymentData)
	if err != nil {
		return fmt.Errorf("failed to update payment: %v", err)
	}

	return nil
}

// Admin operations
func (s *SupabaseService) GetAdminStats(ctx context.Context) (*models.AdminStats, error) {
	// Get total users
	var totalUsers int
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&totalUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to get total users: %v", err)
	}

	// Get total series
	var totalSeries int
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM series").Scan(&totalSeries)
	if err != nil {
		return nil, fmt.Errorf("failed to get total series: %v", err)
	}

	// Get total episodes
	var totalEpisodes int
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM episodes").Scan(&totalEpisodes)
	if err != nil {
		return nil, fmt.Errorf("failed to get total episodes: %v", err)
	}

	// Get total revenue (sum of all completed payments)
	var totalRevenue int
	err = s.db.QueryRowContext(ctx, "SELECT COALESCE(SUM(amount), 0) FROM payments WHERE status = 'completed'").Scan(&totalRevenue)
	if err != nil {
		return nil, fmt.Errorf("failed to get total revenue: %v", err)
	}

	// Get monthly revenue
	var monthlyRevenue int
	err = s.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(amount), 0) FROM payments 
		WHERE status = 'completed' AND created_at >= NOW() - INTERVAL '1 month'
	`).Scan(&monthlyRevenue)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly revenue: %v", err)
	}

	// Get active users (users with activity in last 30 days)
	var activeUsers int
	err = s.db.QueryRowContext(ctx, `
		SELECT COUNT(DISTINCT user_id) FROM purchases 
		WHERE created_at >= NOW() - INTERVAL '30 days'
	`).Scan(&activeUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to get active users: %v", err)
	}

	stats := &models.AdminStats{
		TotalUsers:     totalUsers,
		TotalSeries:    totalSeries,
		TotalEpisodes:  totalEpisodes,
		TotalRevenue:   totalRevenue,
		MonthlyRevenue: monthlyRevenue,
		ActiveUsers:    activeUsers,
	}

	return stats, nil
}

// Helper methods
func (s *SupabaseService) GetClient() interface{} {
	return s.db
}

// Close closes the database connection
func (s *SupabaseService) Close() error {
	return s.db.Close()
}
