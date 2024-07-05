package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Change baseURL to a variable
var baseURL = "https://reqres.in/api"

func main() {
	// Perform a GET request
	users, err := getUsers()
	if err != nil {
		log.Fatalf("Error getting users: %v", err)
	}
	fmt.Println("Users:", string(users))

	// Perform a POST request
	user, err := createUser("John", "Developer")
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Println("Created User:", string(user))
}

func getUsers() ([]byte, error) {
	resp, err := http.Get(baseURL + "/users?page=2")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func createUser(name, job string) ([]byte, error) {
	data := map[string]string{
		"name": name,
		"job":  job,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
