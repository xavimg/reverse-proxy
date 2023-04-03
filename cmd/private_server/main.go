package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Users []User

func dialed(w http.ResponseWriter, r *http.Request) {
	fakeDB := Users{
		User{
			Name: "xavier",
			Type: "M",
		},
		User{
			Name: "marcus",
			Type: "M",
		},
	}

	json, err := json.Marshal(fakeDB)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func main() {
	http.HandleFunc("/", dialed)

	http.ListenAndServe(":8080", nil)
}
