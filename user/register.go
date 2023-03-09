package main

import (
    "fmt"
)

type User struct {
    Username string
    Password string
    Email    string
}

var users []User

func register(username, password, email string) error {
    // Check if user already exists
    for _, user := range users {
        if user.Username == username {
            return fmt.Errorf("username '%s' already exists", username)
        }
    }

    // Create new user
    newUser := User{
        Username: username,
        Password: password,
        Email:    email,
    }
    users = append(users, newUser)
    return nil
}

func main() {
    // Sample user registration
    err := register("johndoe", "password123", "johndoe@example.com")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("User registered successfully!")
}

// In this example code, the User struct defines the properties of a user, which are Username, Password, and Email. The users variable is a slice that holds all the registered users.

// The register function takes three parameters, which are the username, password, and email of the user to be registered. It first checks if the username already exists in the users slice, and returns an error if it does. If the username is unique, a new User instance is created with the given properties and appended to the users slice.

// In the main function, you can see an example of how to register a user by calling the register function with sample values. The function returns an error if the registration fails, or a success message if it is successful.