package main

import (
    "fmt"
    "strings"
)

type Article struct {
    ID          int
    Title       string
    Body        string
    Author      string
    Category    string // new field for category
}

var articles = []Article{}

func createArticle() {
    var title, body, author, category string

    // Prompt the user for article information
    fmt.Println("Create new article")
    fmt.Print("Title: ")
    fmt.Scanln(&title)
    fmt.Print("Body: ")
    fmt.Scanln(&body)
    fmt.Print("Author: ")
    fmt.Scanln(&author)
    fmt.Print("Category: ")
    fmt.Scanln(&category)

    // Create a new article object
    article := Article{
        ID:          len(articles) + 1,
        Title:       title,
        Body:        body,
        Author:      author,
        Category:    category,
    }

    // Add the new article to the list of articles
    articles = append(articles, article)

    fmt.Println("Article created successfully")
}

func searchArticles() {
    searchQuery := ""
    fmt.Println("Search for articles")
    fmt.Print("Enter search query: ")
    fmt.Scanln(&searchQuery)

    // Initialize a map to store the search results for each category
    searchResults := make(map[string][]Article)

    // Search for articles that match the query
    for _, article := range articles {
        if strings.Contains(strings.ToLower(article.Title), strings.ToLower(searchQuery)) || strings.Contains(strings.ToLower(article.Body), strings.ToLower(searchQuery)) {
            // Add the matching article to the search results map for its category
            searchResults[article.Category] = append(searchResults[article.Category], article)
        }
    }

    // Print the search results for each category
    fmt.Printf("\nSearch results for '%s':\n", searchQuery)
    for category, articles := range searchResults {
        fmt.Printf("\n%s:\n", category)
        for _, article := range articles {
            fmt.Printf("\t%s: %s\n", article.Title, article.Body)
        }
    }
}

func main() {
    // Create some sample articles
    articles = []Article{
        {ID: 1, Title: "Introduction to Golang", Body: "This article introduces Golang.", Author: "John Doe", Category: "Programming"},
        {ID: 2, Title: "Golang vs Python", Body: "This article compares Golang and Python.", Author: "Jane Smith", Category: "Programming"},
        {ID: 3, Title: "Family Law in the USA",
