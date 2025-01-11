package repository

import (
	"coding-test-11-01-2025-dating-app/internal/model"

	"gorm.io/gorm"
)

type SwipeRepository struct {
	db *gorm.DB
}

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) CreateProfile(profile *model.Profile) error {
	return r.db.Create(profile).Error
}

func (r *ProfileRepository) GetUnswiped(userID uint, limit int) ([]model.Profile, error) {
	var profiles []model.Profile

	err := r.db.
		Joins("LEFT JOIN swipes ON profiles.user_id = swipes.swiped_id AND swipes.swiper_id = ?", userID).
		Where("swipes.id IS NULL AND profiles.user_id != ?", userID).
		Limit(limit).
		Find(&profiles).Error

	return profiles, err
}
