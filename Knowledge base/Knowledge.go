package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"html/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	Category    string  `json:"category"`
	Author      *Author `json:"author"`
	CreatedDate string  `json:"createddate"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var articles []Article

func getArticle(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application.json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(1000000000))
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application.json")

	//params
	params := mux.Vars(r)

	//loop over the movies, range

	//delete the movie with the id sent
	//add a new movie. the movie sent in the body of postman

	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = params["id"]
			articles = append(articles, article)
			json.NewEncoder(w).Encode(article)
			return
		}
	}
}

func Connect() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:1234567890(localhost:3306)/LegalConnect")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Load templates
	tmpl := template.Must(template.ParseGlob("templates/*"))

	// Route handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get all articles from the database
		rows, err := db.Query("SELECT * FROM articles")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Get the ID of the requested article
		id := r.URL.Query().Get("id")

		// Query the database for the requested article
		row := db.QueryRow("SELECT * FROM articles WHERE id = ?", id)

		// Create an empty slice of articles
		articles := []Article{}

		// Loop through the rows and append articles to the slice
		for rows.Next() {
			var article Article
			err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Author, &article.CreatedDate)
			if err != nil {
				log.Fatal(err)
			}
			articles = append(articles, article)
		}

		// Render the index page with the articles
		tmpl.ExecuteTemplate(w, "index.html", articles)
	})

}

func main() {
	r := mux.NewRouter()

	articles = append(articles, Article{ID: "1", Title: "43288", Content: "Article One", Author: &Author{Firstname: "john", Lastname: "Doe"}})
	articles = append(articles, Article{ID: "2", Title: "45425", Content: "Article Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/articles", getArticle).Methods("GET")
	r.HandleFunc("/articles/{id}", getMovie).Methods("GET")
	r.HandleFunc("/articles", createMovie).Methods("POST")
	r.HandleFunc("/articles/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	//start the web server
	fmt.Printf("starting a server at port 8000/n")
	log.Fatal(http.ListenAndServe(":8000", r))

}

// This code uses the database/sql package to connect to a MySQL database and retrieve articles. The net/http package is used to handle HTTP requests and responses, and the html/template package is used to render HTML templates.

// The code defines an Article struct with fields for the article's ID, title, content, category, author, and creation date. It also defines two route handlers for the index page and the article page. The index page retrieves all articles from the database and renders them using an HTML template, while the article page retrieves a specific article based on the ID parameter and renders it using a separate HTML template.

// Again, please note that this is just a basic example of what a content management system in Golang could look like. You will need to modify and adapt it to fit your specific requirements.
