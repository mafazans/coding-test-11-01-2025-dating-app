package model

import (
	"time"

	"gorm.io/gorm"
)

type SubscriptionStatus string

const (
	SubscriptionActive    SubscriptionStatus = "active"
	SubscriptionExpired   SubscriptionStatus = "expired"
	SubscriptionCancelled SubscriptionStatus = "cancelled"
)

type Subscription struct {
	gorm.Model
	UserID    uint               `gorm:"uniqueIndex" json:"user_id"`
	Status    SubscriptionStatus `json:"status"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date"`
}
