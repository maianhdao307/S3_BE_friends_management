package models

import "gorm.io/gorm"

// Friend is a friend connection between two email addresses
type Friend struct {
	gorm.Model
	Email1 string `json:"email1"`
	Email2 string `json:"email2"`
}
