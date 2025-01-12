package repository

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"context"
)

type RepositoryInterface interface {
	CreateSubscription(ctx context.Context, sub *model.Subscription) error
	GetActiveSubscription(ctx context.Context, userID uint) (*model.Subscription, error)
	IsUserPremium(ctx context.Context, userID uint) (bool, error)
	VerifyProfile(ctx context.Context, userID uint) error
	CreateSwipe(ctx context.Context, swipe *model.Swipe) error
	GetDailySwipeCount(ctx context.Context, userID uint) (int, error)
	CheckMatch(ctx context.Context, userID, profileID uint) (bool, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	Delete(ctx context.Context, user *model.User) error
	CreateProfile(ctx context.Context, profile *model.Profile) error
	GetUnswiped(ctx context.Context, userID uint, limit int) ([]model.Profile, error)
}
