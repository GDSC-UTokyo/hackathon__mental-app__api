package model

import (
	"time"

	"gorm.io/gorm"
)

type ReasonsOnMentalPoints struct {
	Id            string    `json:"id" gorm:"primaryKey"`
	ReasonId      string    `json:"reason_id" gorm:"not null"`
	MentalPointId string    `json:"mental_point_id" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ROMPs []ReasonsOnMentalPoints

type ReasonIdList []string

func (p *ReasonsOnMentalPoints) RegisterReasonsOnMentalPoint() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *ROMPs) RegisterReasonsOnMentalPoints() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *ReasonsOnMentalPoints) GetReportByMentalPointId(mentalPointId string) (tx *gorm.DB) {
	return db.Where("mental_point_id", mentalPointId).Find(&p)
}

func (p *ReasonIdList) GetReasonIdsByMentalPointId(mentalPointId string) (tx *gorm.DB) {
	return db.Model(&ReasonsOnMentalPoints{}).Select("reason_id").Where("mental_point_id", mentalPointId).Find(&p)
}

func (p *ReasonsOnMentalPoints) UpdateReasonsOnMentalPoint() (tx *gorm.DB) {
	return db.Model(&p).Update("reason_id", p.ReasonId)
}

func (p *ROMPs) DeleteReportsByPointId(mentalPointId string) (tx *gorm.DB) {
	return db.Where("mental_point_id = ?", mentalPointId).Unscoped().Delete(&p)
}

func (p *ROMPs) DeleteReportsByReasonId(reasonId string) (tx *gorm.DB) {
	return db.Where("reason_id = ?", reasonId).Unscoped().Delete(&p)
}
