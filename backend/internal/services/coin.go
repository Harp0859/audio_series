package services

import (
	"context"
	"fmt"

	"audio-series-app/backend/internal/models"

	"github.com/google/uuid"
)

type CoinService struct {
	supabase *SupabaseService
}

func NewCoinService(supabase *SupabaseService) *CoinService {
	return &CoinService{
		supabase: supabase,
	}
}

func (s *CoinService) UnlockEpisode(ctx context.Context, userIDStr, episodeIDStr string) error {
	// Parse UUIDs
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	episodeID, err := uuid.Parse(episodeIDStr)
	if err != nil {
		return fmt.Errorf("invalid episode ID: %w", err)
	}

	// Get episode to check price
	episode, err := s.supabase.GetEpisodeByID(ctx, episodeID)
	if err != nil {
		return fmt.Errorf("failed to get episode: %w", err)
	}

	// Check if user already owns the episode
	isOwned, err := s.supabase.HasUserPurchasedEpisode(ctx, userID, episodeID)
	if err != nil {
		return fmt.Errorf("failed to check purchase status: %w", err)
	}
	if isOwned {
		return fmt.Errorf("episode already owned")
	}

	// Get user to check coin balance
	user, err := s.supabase.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Check if user has enough coins
	if user.CoinBalance < episode.CoinPrice {
		return fmt.Errorf("insufficient coins")
	}

	// Deduct coins from user balance
	err = s.supabase.UpdateUserCoins(ctx, userID, -episode.CoinPrice)
	if err != nil {
		return fmt.Errorf("failed to update coin balance: %w", err)
	}

	// Create purchase record
	purchase := &models.Purchase{
		UserID:    userID,
		EpisodeID: &episodeID,
		Type:      "episode",
		Amount:    episode.CoinPrice,
		Status:    "completed",
	}

	err = s.supabase.CreatePurchase(ctx, purchase)
	if err != nil {
		return fmt.Errorf("failed to create purchase record: %w", err)
	}

	// Create coin transaction record
	referenceID := purchase.ID.String()
	transaction := &models.CoinTransaction{
		UserID:      userID,
		Type:        "purchase",
		Amount:      -episode.CoinPrice,
		Balance:     user.CoinBalance - episode.CoinPrice,
		Description: fmt.Sprintf("Purchased episode: %s", episode.Title),
		ReferenceID: &referenceID,
	}

	err = s.supabase.CreateCoinTransaction(ctx, transaction)
	if err != nil {
		return fmt.Errorf("failed to create coin transaction: %w", err)
	}

	return nil
}

func (s *CoinService) UnlockSeries(ctx context.Context, userIDStr, seriesIDStr string) error {
	// Parse UUIDs
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	seriesID, err := uuid.Parse(seriesIDStr)
	if err != nil {
		return fmt.Errorf("invalid series ID: %w", err)
	}

	// Get all episodes in the series
	episodes, err := s.supabase.GetEpisodesBySeriesID(ctx, seriesID)
	if err != nil {
		return fmt.Errorf("failed to get series episodes: %w", err)
	}

	// Calculate total cost
	totalCost := 0
	for _, episode := range episodes {
		isOwned, err := s.supabase.HasUserPurchasedEpisode(ctx, userID, episode.ID)
		if err != nil {
			return fmt.Errorf("failed to check episode ownership: %w", err)
		}
		if !isOwned {
			totalCost += episode.CoinPrice
		}
	}

	if totalCost == 0 {
		return fmt.Errorf("all episodes already owned")
	}

	// Get user to check coin balance
	user, err := s.supabase.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Check if user has enough coins
	if user.CoinBalance < totalCost {
		return fmt.Errorf("insufficient coins")
	}

	// Deduct coins from user balance
	err = s.supabase.UpdateUserCoins(ctx, userID, -totalCost)
	if err != nil {
		return fmt.Errorf("failed to update coin balance: %w", err)
	}

	// Create purchase records for each episode
	for _, episode := range episodes {
		isOwned, err := s.supabase.HasUserPurchasedEpisode(ctx, userID, episode.ID)
		if err != nil {
			return fmt.Errorf("failed to check episode ownership: %w", err)
		}
		if !isOwned {
			purchase := &models.Purchase{
				UserID:    userID,
				EpisodeID: &episode.ID,
				Type:      "episode",
				Amount:    episode.CoinPrice,
				Status:    "completed",
			}

			err = s.supabase.CreatePurchase(ctx, purchase)
			if err != nil {
				return fmt.Errorf("failed to create purchase record: %w", err)
			}
		}
	}

	// Create coin transaction record
	transaction := &models.CoinTransaction{
		UserID:      userID,
		Type:        "purchase",
		Amount:      -totalCost,
		Balance:     user.CoinBalance - totalCost,
		Description: "Purchased entire series",
	}

	err = s.supabase.CreateCoinTransaction(ctx, transaction)
	if err != nil {
		return fmt.Errorf("failed to create coin transaction: %w", err)
	}

	return nil
}

func (s *CoinService) AddCoins(ctx context.Context, userID uuid.UUID, amount int, description string) error {
	// Add coins to user balance
	err := s.supabase.UpdateUserCoins(ctx, userID, amount)
	if err != nil {
		return fmt.Errorf("failed to update coin balance: %w", err)
	}

	// Get updated user to get new balance
	user, err := s.supabase.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Create coin transaction record
	transaction := &models.CoinTransaction{
		UserID:      userID,
		Type:        "admin",
		Amount:      amount,
		Balance:     user.CoinBalance,
		Description: description,
	}

	err = s.supabase.CreateCoinTransaction(ctx, transaction)
	if err != nil {
		return fmt.Errorf("failed to create coin transaction: %w", err)
	}

	return nil
}
