package repository

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"context"
	"time"

	"gorm.io/gorm"
)

func (r *Repository) CreateSubscription(ctx context.Context, sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *Repository) GetActiveSubscription(ctx context.Context, userID uint) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.Where("user_id = ? AND status = ? AND end_date > ?",
		userID, model.SubscriptionActive, time.Now()).
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *Repository) IsUserPremium(ctx context.Context, userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Subscription{}).
		Where("user_id = ? AND status = ? AND end_date > ?",
			userID, model.SubscriptionActive, time.Now()).
		Count(&count).Error
	return count > 0, err
}

// Update Repository to include verification
func (r *Repository) VerifyProfile(ctx context.Context, userID uint) error {
	return r.db.Model(&model.Profile{}).
		Where("user_id = ?", userID).
		Update("is_verified", true).Error
}

func (r *Repository) CreateSwipe(ctx context.Context, swipe *model.Swipe) error {
	// Todo: handle validation if unique constraint error
	return r.db.Create(swipe).Error
}

func (r *Repository) GetDailySwipeCount(ctx context.Context, userID uint) (int, error) {
	var count int64
	startOfDay := time.Now().UTC().Truncate(24 * time.Hour)

	err := r.db.Model(&model.Swipe{}).
		Where("swiper_id = ? AND created_at >= ?", userID, startOfDay).
		Count(&count).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return int(count), nil
}

// CheckMatch checks if there's a mutual like between users
func (r *Repository) CheckMatch(ctx context.Context, userID, profileID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Swipe{}).
		Where("swiper_id = ? AND swiped_id = ? AND is_like = ? AND created_at <= NOW()",
			profileID, userID, true).
		Count(&count).Error

	return count > 0, err
}

func (r *Repository) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Delete(ctx context.Context, user *model.User) error {
	return r.db.Delete(user).Error
}

func (r *Repository) CreateProfile(ctx context.Context, profile *model.Profile) error {
	return r.db.Create(profile).Error
}

func (r *Repository) GetUnswiped(ctx context.Context, userID uint, limit int) ([]model.Profile, error) {
	var profiles []model.Profile

	err := r.db.
		Joins("LEFT JOIN swipes ON profiles.user_id = swipes.swiped_id AND swipes.swiper_id = ?", userID).
		Where("swipes.id IS NULL AND profiles.user_id != ?", userID).
		Limit(limit).
		Find(&profiles).Error

	return profiles, err
}
