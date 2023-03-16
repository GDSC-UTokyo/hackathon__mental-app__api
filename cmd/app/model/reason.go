package model

import (
	"gorm.io/gorm"
	"time"
)

type Reason struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	Reason    string         `json:"reason" gorm:"unique;not null"`
	UserId    string         `json:"user_id" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Reasons []Reason

func (p *Reasons) GetReasonsByUserId(userId string) (tx *gorm.DB) {
	return db.Where("user_id=", userId).Find(&p)
}
