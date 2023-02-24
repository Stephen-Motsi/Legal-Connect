package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	ZipCode        string   `json:"zip_code"`
	LegalInterests []string `json:"legal_interests"`
}

var users []User

func registerUser(w http.ResponseWriter, r *http.Request) {
	// Decode the request body to a User struct
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email is already taken
	for _, user := range users {
		if user.Email == newUser.Email {
			http.Error(w, "Email already taken", http.StatusConflict)
			return
		}
	}

	// Add the user to the users slice
	users = append(users, newUser)

	// Redirect the user to the dashboard or another page after successful registration
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

	// Send a response indicating success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

func main() {
	// Create a new HTTP server and register the registerUser handler function
	http.HandleFunc("/users", registerUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
