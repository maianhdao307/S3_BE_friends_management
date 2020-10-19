package models

import "gorm.io/gorm"

// Friend is a friend connection between two email addresses
type Friend struct {
	gorm.Model
	Email1 string `json:"email1" gorm:"UNIQUE_INDEX:compositeindex"`
	Email2 string `json:"email2" gorm:"UNIQUE_INDEX:compositeindex"`
}
