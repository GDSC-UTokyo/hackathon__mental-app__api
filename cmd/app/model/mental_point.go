package model

import (
	"time"

	"gorm.io/gorm"
)

type MentalPoint struct {
	Id          string         `json:"id" gorm:"primaryKey"`
	Point       int            `json:"point" gorm:"not null"`
	UserId      string         `json:"user_id" gorm:"not null"`
	CreatedDate string         `json:"created_date" gorm:"not null;size:255"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type MentalPoints []MentalPoint

func (p *MentalPoint) RegisterPoint() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *MentalPoint) GetReportByMentalPointId(mentalPointId string) (tx *gorm.DB) {
	return db.Where("id", mentalPointId).Find(&p)
}

func (p *MentalPoints) GetReportsByDate(userId string, startDate string, endDateNext string) (tx *gorm.DB) {
	return db.Where("user_id", userId).Where("created_date >= ?", startDate).Where("created_date < ?", endDateNext).Order("created_date desc").Find(&p)
}

func (p *MentalPoints) GetReportsByCount(userId string, count int) (tx *gorm.DB) {
	return db.Where("user_id", userId).Order("created_date desc").Limit(count).Find(&p)
}

func (p *MentalPoints) GetReportsByDateAndCount(userId string, startDate string, endDateNext string, count int) (tx *gorm.DB) {
	return db.Where("user_id", userId).Where("created_date >= ?", startDate).Where("created_date < ?", endDateNext).Order("created_date desc").Limit(count).Find(&p)
}

func (p *MentalPoint) UpdateMentalPoint() (tx *gorm.DB) {
	return db.Model(&p).Update("point", p.Point)
}

func (p *MentalPoint) GetReportByUserIdAndDate(userId string, createdDate string) (tx *gorm.DB) {
	return db.Where("user_id", userId).Where("created_date", createdDate).Find(&p)
}
