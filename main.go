package main

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
	"./models"
	"github.com/go-chi/chi"
)

type rootHandler func(w http.ResponseWriter, r *http.Request) error

// ClientError is an error whose details to be shared with client.
type ClientError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

func (fn rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r) // Call handler function
	if err == nil {
		return
	}
	// This is where our error handling logic starts.
	log.Printf("An error accured: %v", err) // Log the error.

	clientError, ok := err.(ClientError) // Check if it is a ClientError.
	if !ok {
		// If the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(500) // return 500 Internal Server Error.
		return
	}

	body, err := clientError.ResponseBody() // Try to get response body of ClientError.
	if err != nil {
		log.Printf("An error accured: %v", err)
		w.WriteHeader(500)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}

func main() {
	err := models.Connect()
	if err != nil {
		fmt.Println("Connect to database failed")
	}

	r := chi.NewRouter()
	r.Route("/friends", func(r chi.Router) {
		r.Method("post", "/", rootHandler(handler.CreateFriend))
		r.Method("get", "/friends-list/{email}", rootHandler(handler.GetFriendsListOfEmail))
		r.Method("get", "/common-friends", rootHandler(handler.GetCommonFriends))
	})
	r.Route("/subscriptions", func(r chi.Router) {
		r.Method("post", "/", rootHandler(handler.CreateSubscription))
		// r.Method("get", "/emails-receive-updates", rootHandler(handler.GetEmailsReceiveUpdates))
	})
	r.Route("/blockings", func(r chi.Router) {
		r.Method("post", "/", rootHandler(handler.CreateBlocking))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
