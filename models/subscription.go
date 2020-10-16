package models

import "gorm.io/gorm"

// Subscription is to subscribe to updates from an email address
type Subscription struct {
	gorm.Model
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
