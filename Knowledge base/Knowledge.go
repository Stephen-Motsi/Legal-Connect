package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    // "html/template"
    _"github.com/go-sql-driver/mysql"
)

type Article struct {
    ID          int
    Title       string
    Content     string
    Category    string
    Author      string
    CreatedDate string
}

func main() {
    // Connect to the database
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/knowledgebase")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Load templates
    // tmpl := template.Must(template.ParseGlob("templates/*"))

    // Route handlers
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Get all articles from the database
        rows, err := db.Query("SELECT * FROM articles")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        // Create an empty slice of articles
        articles := []Article{}

        // Loop through the rows and append articles to the slice
        for rows.Next() {
            var article Article
            err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Author, &article.CreatedDate)
            if err != nil {
                log.Fatal(err)
            }
            articles = append(articles, article)
        }

        // Render the index page with the articles
        // tmpl.ExecuteTemplate(w, "index.html", articles)
    })

    http.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
        // Get the ID of the requested article
        id := r.URL.Query().Get("id")

        // Query the database for the requested article
        row := db.QueryRow("SELECT * FROM articles WHERE id = ?", id)

        // Create an empty article
        var article Article

        // Scan the row into the article
        err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Author, &article.CreatedDate)
        if err != nil {
            log.Fatal(err)
        }

        // Render the article page with the article
        // tmpl.ExecuteTemplate(w, "article.html", article)
    })

    // Start the web server
    fmt.Println("Listening on :3000...")
    http.ListenAndServe(":3000", nil)
}









// This code uses the database/sql package to connect to a MySQL database and retrieve articles. The net/http package is used to handle HTTP requests and responses, and the html/template package is used to render HTML templates.

// The code defines an Article struct with fields for the article's ID, title, content, category, author, and creation date. It also defines two route handlers for the index page and the article page. The index page retrieves all articles from the database and renders them using an HTML template, while the article page retrieves a specific article based on the ID parameter and renders it using a separate HTML template.

// Again, please note that this is just a basic example of what a content management system in Golang could look like. You will need to modify and adapt it to fit your specific requirements.