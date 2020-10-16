package main

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
	"./models"
	"github.com/go-chi/chi"
)

func main() {
	err := models.Connect()
	if err != nil {
		fmt.Println("Connect to database failed")
	}

	r := chi.NewRouter()
	r.Route("/friend", func(r chi.Router) {
		r.Post("/", handler.CreateFriend)
		r.Get("/", handler.GetFriends)
		r.Get("/{email}", handler.GetFriendsListOfEmail)
		r.Get("/common-friends", handler.GetCommonFriends)
	})
	r.Route("/subscription", func(r chi.Router) {
		r.Post("/", handler.CreateSubscription)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
