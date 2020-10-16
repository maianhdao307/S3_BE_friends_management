package handler

import (
	"encoding/json"
	"net/http"

	"../models"
	"../utils"
	"github.com/go-chi/chi"
)

// CreateFriend creates a friend connection between two email addresses
func CreateFriend(w http.ResponseWriter, r *http.Request) {
	var friend models.Friend
	json.NewDecoder(r.Body).Decode(&friend)
	models.DB.Create(&friend)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

// GetFriends retrieves all friends list
func GetFriends(w http.ResponseWriter, r *http.Request) {
	var friends []models.Friend
	models.DB.Find(&friends)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"friends": friends,
		"count":   len(friends),
	})
}

// GetFriendsListOfEmail retrieves the friends list for an email address
func GetFriendsListOfEmail(w http.ResponseWriter, r *http.Request) {
	var friends []models.Friend
	email := chi.URLParam(r, "email")
	models.DB.Where("email1 = ?", email).Or("email2 = ?", email).Find(&friends)
	emails := make([]string, len(friends))
	for i, friend := range friends {
		if friend.Email1 == email {
			emails[i] = friend.Email2
		} else {
			emails[i] = friend.Email1
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"friends": emails,
		"count":   len(friends),
	})
}

func getFriendsOfEmail(email string) []string {
	var friends []models.Friend
	models.DB.Where("email1 = ?", email).Or("email2 = ?", email).Find(&friends)
	emails := make([]string, len(friends))
	for i, friend := range friends {
		if friend.Email1 == email {
			emails[i] = friend.Email2
		} else {
			emails[i] = friend.Email1
		}
	}
	return emails
}

type friendsRequest struct {
	Friends []string `json:"friends"`
}

// GetCommonFriends retrieves the common friends list between two email addresses
func GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	var body friendsRequest
	json.NewDecoder(r.Body).Decode(&body)
	emails1 := getFriendsOfEmail(body.Friends[0])
	emails2 := getFriendsOfEmail(body.Friends[1])

	commonEmails := utils.Intersection(emails1, emails2)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"friends": commonEmails,
		"count":   len(commonEmails),
	})
}
