package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"../models"
	"../utils"
	"github.com/go-chi/chi"
)

// CreateFriend creates a friend connection between two email addresses
func CreateFriend(w http.ResponseWriter, r *http.Request) error {
	var friend models.Friend
	json.NewDecoder(r.Body).Decode(&friend)
	var blocking models.Blocking
	err := models.DB.Where("requestor = ? AND target = ?", friend.Email1, friend.Email2).Or("requestor = ? AND target = ?", friend.Email2, friend.Email1).First(&blocking).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return NewHTTPError(nil, http.StatusPreconditionFailed, "An email address is blocked by the other email address")
	}

	err = models.DB.Create(&friend).Error
	if err != nil {
		return fmt.Errorf("Create friend failed: %v", err)
	}

	json.NewEncoder(w).Encode(map[string]bool{"success": true})
	return nil
}

func getFriendsOfEmail(email string) ([]string, error) {
	var friends []models.Friend
	err := models.DB.Where("email1 = ?", email).Or("email2 = ?", email).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	emails := make([]string, len(friends))
	for i, friend := range friends {
		if friend.Email1 == email {
			emails[i] = friend.Email2
		} else {
			emails[i] = friend.Email1
		}
	}
	return emails, nil
}

// GetFriendsListOfEmail retrieves the friends list for an email address
func GetFriendsListOfEmail(w http.ResponseWriter, r *http.Request) error {
	email := chi.URLParam(r, "email")

	emails, err := getFriendsOfEmail(email)
	if err != nil {
		return fmt.Errorf("Get friends failed %v", err)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"friends": emails,
		"count":   len(emails),
	})
	return nil
}

// GetCommonFriends retrieves the common friends list between two email addresses
func GetCommonFriends(w http.ResponseWriter, r *http.Request) error {
	type requestBody struct {
		Friends []string `json:"friends"`
	}
	var body requestBody
	json.NewDecoder(r.Body).Decode(&body)
	var err error
	var emails1, emails2 []string
	emails1, err = getFriendsOfEmail(body.Friends[0])
	emails2, err = getFriendsOfEmail(body.Friends[1])

	if err != nil {
		return fmt.Errorf("Get friends failed %v", err)
	}

	commonEmails := utils.Intersection(emails1, emails2)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"friends": commonEmails,
		"count":   len(commonEmails),
	})
	return nil
}
