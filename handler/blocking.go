package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
)

// CreateBlocking creates a blocking to block updates from an email address
func CreateBlocking(w http.ResponseWriter, r *http.Request) error {
	var blocking models.Blocking
	json.NewDecoder(r.Body).Decode(&blocking)

	err := models.DB.Create(&blocking).Error
	if err != nil {
		return fmt.Errorf("Create blocking failed: %v", err)
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
	return nil
}
