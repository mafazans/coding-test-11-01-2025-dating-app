package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID     uint   `gorm:"uniqueIndex" json:"user_id"`
	Name       string `json:"name"`
	Bio        string `json:"bio"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	PhotoURL   string `json:"photo_url"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
}
