package repository

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"time"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) CreateSubscription(sub *model.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *SubscriptionRepository) GetActiveSubscription(userID uint) (*model.Subscription, error) {
	var sub model.Subscription
	err := r.db.Where("user_id = ? AND status = ? AND end_date > ?",
		userID, model.SubscriptionActive, time.Now()).
		First(&sub).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) IsUserPremium(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Subscription{}).
		Where("user_id = ? AND status = ? AND end_date > ?",
			userID, model.SubscriptionActive, time.Now()).
		Count(&count).Error
	return count > 0, err
}

// Update ProfileRepository to include verification
func (r *ProfileRepository) VerifyProfile(userID uint) error {
	return r.db.Model(&model.Profile{}).
		Where("user_id = ?", userID).
		Update("is_verified", true).Error
}
