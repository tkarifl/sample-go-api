package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Phone string `gorm:"not null" json:"phone"`
	Mail  string `gorm:"not null" json:"mail"`
}
