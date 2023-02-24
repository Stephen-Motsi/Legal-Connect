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

	// Send a response indicating success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")

	// Redirect the user to the dashboard or another page after successful registration
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	// Decode the request body to a User struct
	var loginUser User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user with the given email and password exists
	for _, user := range users {
		if user.Email == loginUser.Email && user.Password == loginUser.Password {
			// Redirect the user to the dashboard or another page after successful login
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	}

	// If no matching user is found, send an HTTP response indicating failure
	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
}

func main() {
	// Create a new HTTP server and register the registerUser and loginUser handler functions
	http.HandleFunc("/register", registerUser)
	http.HandleFunc("/login", loginUser)

	// Serve static files such as HTML, CSS, and JavaScript from the "homepage" directory
	fs := http.FileServer(http.Dir("homepage"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
