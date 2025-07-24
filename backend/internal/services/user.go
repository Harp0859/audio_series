package services

import (
	"context"
	"fmt"

	"audio-series-app/backend/internal/models"

	"github.com/google/uuid"
)

type UserService struct {
	supabase *SupabaseService
}

func NewUserService(supabase *SupabaseService) *UserService {
	return &UserService{
		supabase: supabase,
	}
}

func (s *UserService) GetUserProfile(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return s.supabase.GetUserByID(ctx, userID)
}

func (s *UserService) UpdateUserProfile(ctx context.Context, user *models.User) error {
	return s.supabase.UpdateUser(ctx, user)
}

func (s *UserService) GetUserPurchases(ctx context.Context, userID uuid.UUID) ([]*models.Purchase, error) {
	return s.supabase.GetUserPurchases(ctx, userID)
}

func (s *UserService) GetUserCoinBalance(ctx context.Context, userID uuid.UUID) (int, error) {
	user, err := s.supabase.GetUserByID(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to get user: %w", err)
	}
	return user.CoinBalance, nil
}
