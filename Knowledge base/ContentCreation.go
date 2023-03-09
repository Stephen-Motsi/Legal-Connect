package main

import (
    "fmt"
    "time"
)

// Article represents a single article in the knowledge base
type Article struct {
    Title       string
    Description string
    Category    string
    CreatedAt   time.Time
}

// createNewArticle creates a new Article with the specified title, description, and category
func createNewArticle(title, description, category string) Article {
    return Article{
        Title:       title,
        Description: description,
        Category:    category,
        CreatedAt:   time.Now(),
    }
}

func main() {
    // Initialize an empty slice of Articles
    articles := []Article{}

    // Create a new Article and add it to the slice
    newArticle := createNewArticle("How to Write a Will", "A guide to writing a will for beginners", "Estate Planning")
    articles = append(articles, newArticle)

    // Print out the newly created article
    fmt.Println("New article created:")
    fmt.Printf("\tTitle: %s\n", newArticle.Title)
    fmt.Printf("\tDescription: %s\n", newArticle.Description)
    fmt.Printf("\tCategory: %s\n", newArticle.Category)
    fmt.Printf("\tCreated At: %s\n", newArticle.CreatedAt)
}



// In this example, we define the Article struct and a createNewArticle function to create a new Article. The function takes in the title, description, and category as arguments, and sets the CreatedAt field to the current time using time.Now().

// In the main function, we initialize an empty slice of Articles and create a new article using the createNewArticle function. We then append the new article to the articles slice and print out the details of the new article using fmt.Printf() statements.

// Note that this is just one example of how to create new articles, and you can modify the code to fit your specific needs, for example by getting the article details from user input or a form.g