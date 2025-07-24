package services

import (
	"context"

	"audio-series-app/backend/internal/models"

	"github.com/google/uuid"
)

type SeriesService struct {
	supabase *SupabaseService
}

func NewSeriesService(supabase *SupabaseService) *SeriesService {
	return &SeriesService{
		supabase: supabase,
	}
}

func (s *SeriesService) CreateSeries(ctx context.Context, series *models.Series) error {
	return s.supabase.CreateSeries(ctx, series)
}

func (s *SeriesService) GetSeries(ctx context.Context) ([]*models.Series, error) {
	return s.supabase.GetSeries(ctx)
}

func (s *SeriesService) GetSeriesByID(ctx context.Context, seriesID uuid.UUID) (*models.Series, error) {
	return s.supabase.GetSeriesByID(ctx, seriesID)
}

func (s *SeriesService) GetSeriesWithEpisodes(ctx context.Context, seriesID uuid.UUID) (*models.SeriesWithEpisodes, error) {
	series, err := s.supabase.GetSeriesByID(ctx, seriesID)
	if err != nil {
		return nil, err
	}

	episodes, err := s.supabase.GetEpisodesBySeriesID(ctx, seriesID)
	if err != nil {
		return nil, err
	}

	return &models.SeriesWithEpisodes{
		Series:   series,
		Episodes: episodes,
	}, nil
}
