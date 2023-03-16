package model

import (
	"gorm.io/gorm"
	"time"
)

type MentalPoint struct {
	Id          string         `json:"id" gorm:"primaryKey"`
	Point       int            `json:"point" gorm:"not null"`
	UserId      string         `json:"user_id" gorm:"not null"`
	CreatedDate string         `json:"created_date" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
