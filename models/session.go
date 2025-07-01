package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Username     string
	RefreshToken string
}
