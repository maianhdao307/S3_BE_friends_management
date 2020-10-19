package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

// CreateSubscription creates a subscription to updates from an email address
func CreateSubscription(w http.ResponseWriter, r *http.Request) error {
	var subscription models.Subscription
	json.NewDecoder(r.Body).Decode(&subscription)

	err := models.DB.Create(&subscription).Error
	if err != nil {
		return fmt.Errorf("Create subscription failed: %v", err)
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
	return nil
}

// func GetEmailsReceiveUpdates(w http.ResponseWriter, r *http.Request) error {
// 	type requestBody struct {
// 		Sender string `json:"sender"`
// 		Text   string `json:"text"`
// 	}
// 	var body requestBody
// 	json.NewDecoder(r.Body).Decode(&body)

// }
