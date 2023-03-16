package model

import (
	"gorm.io/gorm"
	"time"
)

type ReasonsOnMentalPoints struct {
	Id            string         `json:"id" gorm:"primaryKey"`
	ReasonId      string         `json:"reason_id" gorm:"not null"`
	MentalPointId string         `json:"mental_point_id" gorm:"not null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
