package models

import "gorm.io/gorm"

type InputAuth struct {
	gorm.Model
	Username string
	Password string
	Refresh  string
}
