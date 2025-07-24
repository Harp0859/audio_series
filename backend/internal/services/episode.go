package services

import (
	"context"
	"fmt"

	"audio-series-app/backend/internal/models"

	"github.com/google/uuid"
)

type EpisodeService struct {
	supabase *SupabaseService
}

func NewEpisodeService(supabase *SupabaseService) *EpisodeService {
	return &EpisodeService{
		supabase: supabase,
	}
}

func (s *EpisodeService) CreateEpisode(ctx context.Context, episode *models.Episode) error {
	return s.supabase.CreateEpisode(ctx, episode)
}

func (s *EpisodeService) GetEpisodeByID(ctx context.Context, episodeID uuid.UUID) (*models.Episode, error) {
	return s.supabase.GetEpisodeByID(ctx, episodeID)
}

func (s *EpisodeService) GetEpisodesBySeriesID(ctx context.Context, seriesID uuid.UUID) ([]*models.Episode, error) {
	return s.supabase.GetEpisodesBySeriesID(ctx, seriesID)
}

func (s *EpisodeService) GetEpisodeWithPurchaseStatus(ctx context.Context, episodeIDStr, userIDStr string) (*models.EpisodeWithPurchase, error) {
	// Parse UUIDs
	episodeID, err := uuid.Parse(episodeIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid episode ID: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	episode, err := s.supabase.GetEpisodeByID(ctx, episodeID)
	if err != nil {
		return nil, err
	}

	isOwned, err := s.supabase.HasUserPurchasedEpisode(ctx, userID, episodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to check purchase status: %w", err)
	}

	// Check if user can unlock (has enough coins)
	user, err := s.supabase.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	canUnlock := !isOwned && user.CoinBalance >= episode.CoinPrice

	return &models.EpisodeWithPurchase{
		Episode:   episode,
		IsOwned:   isOwned,
		CanUnlock: canUnlock,
	}, nil
}
