package models

import "gorm.io/gorm"

// Blocking is to block updates from an email address
type Blocking struct {
	gorm.Model
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
