package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	users, _ := GetUser("localhost:8080")

	usersJSON, _ := json.Marshal(users)
	os.Stdout.Write(usersJSON)

}

func GetUser(host string) ([]User, error) {
	uri := fmt.Sprintf("http://%s/users", host)

	resp, err := http.Get(uri)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	var users []User

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
