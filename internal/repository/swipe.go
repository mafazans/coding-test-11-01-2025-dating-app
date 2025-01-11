package repository

import (
	"coding-test-11-01-2025-dating-app/internal/model"
	"time"

	"gorm.io/gorm"
)

func NewSwipeRepository(db *gorm.DB) *SwipeRepository {
	return &SwipeRepository{db: db}
}

func (r *SwipeRepository) CreateSwipe(swipe *model.Swipe) error {
	return r.db.Create(swipe).Error
}

func (r *SwipeRepository) GetDailySwipeCount(userID uint) (int, error) {
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
func (r *SwipeRepository) CheckMatch(userID, profileID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Swipe{}).
		Where("swiper_id = ? AND swiped_id = ? AND is_like = ? AND created_at <= NOW()",
			profileID, userID, true).
		Count(&count).Error

	return count > 0, err
}
