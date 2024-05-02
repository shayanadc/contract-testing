package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
}

func main() {
	user, err := GetUsers("localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func GetUsers(host string) (*[]User, error) {
	uri := fmt.Sprintf("http://%s/users", host)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return &users, nil

}
