package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

// CreateSubscription creates a subscription to updates from an email address
func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var subscription models.Subscription
	json.NewDecoder(r.Body).Decode(&subscription)
	fmt.Println(subscription)
	models.DB.Create(&subscription)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
