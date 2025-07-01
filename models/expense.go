package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	Name   string
	Price  int
	Date   time.Time
	Author string
}
