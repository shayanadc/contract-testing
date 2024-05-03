package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	http.HandleFunc("/users", GetUsers)

	http.ListenAndServe(":8080", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{
			ID:   1,
			Name: "Arjun",
		},
		{
			ID:   2,
			Name: "Sarah",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
