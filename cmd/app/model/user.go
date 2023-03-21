package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique;not null"`
	UId       string         `json:"uid" gorm:"unique;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Users []User

func (p *Users) CreateUsers() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *User) CreateUser() (tx *gorm.DB) {
	return db.Create(&p)
}

func (p *User) GetUserByUId(uid string) (tx *gorm.DB) {
	return db.Where("uid = ?", uid).First(&p)
}
