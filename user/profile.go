package main
import (
    "net/http"
    "encoding/json"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
)
// Define the User struct with necessary fields
type User struct {
    ID       int
    Name     string
    Email    string
    Password string
    ImageURL string
}

// Update user profile information
func updateUserProfile(userID int, name string, email string, password string, imageURL string) error {
    // Connect to the database
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        return err
    }
    defer db.Close()

    // Prepare the SQL statement to update user profile
    stmt, err := db.Prepare("UPDATE users SET name=?, email=?, password=?, image_url=? WHERE id=?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the SQL statement to update user profile
    _, err = stmt.Exec(name, email, password, imageURL, userID)
    if err != nil {
        return err
    }

    return nil
}
