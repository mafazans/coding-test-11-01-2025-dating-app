package model

import (
	"time"

	"gorm.io/gorm"
)

type Swipe struct {
	gorm.Model
	SwiperID uint `gorm:"index" json:"swiper_id"`
	SwipedID uint `json:"swiped_id"`
	IsLike   bool `json:"is_like"` // true for right swipe (like), false for left swipe (dislike)
}

type SwipeCount struct {
	Count     int
	LastReset time.Time
}
