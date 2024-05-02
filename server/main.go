package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
}

func main() {
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8080", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{
			ID:        1,
			FirstName: "Arjun",
		},
		{
			ID:        2,
			FirstName: "Sarah",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
