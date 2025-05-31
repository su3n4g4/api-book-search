package entity

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	memo      string `gorm:"not null"`
	volume_id uint   `gorm:"not null"`
}
