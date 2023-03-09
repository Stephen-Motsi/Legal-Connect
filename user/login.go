package main

import (
	"fmt"
)

// User represents a user account
type User struct {
	Username string
	Password string
}

var users []User

func main() {
	// Create sample user account
	users = []User{
		User{Username: "johndoe", Password: "password123"},
	}

	// Prompt for login credentials
	var username, password string
	fmt.Println("Please enter your login credentials.")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	// Authenticate user
	if isValidUser(username, password) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Invalid login credentials. Please try again.")
	}

	// Password reset
	fmt.Println("Forgot your password? Reset it now.")
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)
	if resetPassword(username) {
		fmt.Println("Password reset successful! Please check your email for instructions.")
	} else {
		fmt.Println("Unable to reset password. Please check your username and try again.")
	}
}

// isValidUser checks if the user credentials are valid
func isValidUser(username, password string) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// resetPassword resets the user's password
func resetPassword(username string) bool {
	for i, user := range users {
		if user.Username == username {
			// Simulate sending password reset email
			fmt.Printf("Sending password reset email to %s...\n", user.Username)
			// Update password in user slice
			users[i].Password = "newpassword123"
			return true
		}
	}
	return false
}
