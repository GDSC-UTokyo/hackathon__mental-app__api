package model

import (
	"time"

	"gorm.io/gorm"
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

func (p *MentalPoint) RegisterPoint() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *MentalPoint) GetReportByMentalPointId(mentalPointId string) (tx *gorm.DB) {
	return db.Where("id", mentalPointId).Find(&p)
}

func (p *MentalPoint) UpdateMentalPoint() (tx *gorm.DB) {
	return db.Model(&p).Update("point", p.Point)
}
