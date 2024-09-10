package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name       string `gorm:"size:100"`
	IsPrimary  bool
	CityID     string `gorm:"size:255"`
	ProvinceID string `gorm:"size:255"`
	Address1   string `gorm:"size:255"`
	Address2   string `gorm:"size:255"`
	Phone      string `gorm:"size:100"`
	Email      string `gorm:"size:100"`
	PostCode   string `gorm:"size:100"`
	User       User
	UserID     string `gorm:"size:36;index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
