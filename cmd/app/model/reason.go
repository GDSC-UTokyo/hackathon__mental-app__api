package model

import (
	"time"

	"gorm.io/gorm"
)

type Reason struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	Reason    string         `json:"reason" gorm:"not null;size:255" `
	UserId    string         `json:"user_id" gorm:"not null;size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Reasons []Reason

func (p *Reasons) GetReasonsByUserId(userId string) (tx *gorm.DB) {
	return db.Where("user_id", userId).Find(&p)
}

func (p *Reason) CreateReason() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *Reasons) CreateReasons() (tx *gorm.DB) {
	return db.Create(&p)
}
