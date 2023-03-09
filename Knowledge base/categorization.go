package main

import (
	"fmt"
	"strings"
)

type Article struct {
	Title       string
	Description string
	Category    string
}

type KnowledgeBase struct {
	Name     string
	Articles []Article
}

func (kb *KnowledgeBase) AddArticle(article Article) {
	kb.Articles = append(kb.Articles, article)
}

func (kb *KnowledgeBase) GetArticlesByCategory(category string) []Article {
	var result []Article
	for _, article := range kb.Articles {
		if strings.ToLower(article.Category) == strings.ToLower(category) {
			result = append(result, article)
		}
	}
	return result
}

func (kb *KnowledgeBase) SearchArticles(query string) []Article {
	var result []Article
	for _, article := range kb.Articles {
		if strings.Contains(strings.ToLower(article.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(article.Description), strings.ToLower(query)) {
			result = append(result, article)
		}
	}
	return result
}

func main() {
	kb := &KnowledgeBase{}

	article1 := Article{
		Title:       "What are my rights if I am arrested?",
		Description: "A brief overview of what to expect when you are arrested.",
		Category:    "Criminal Law",
	}
	kb.AddArticle(article1)

	article2 := Article{
		Title:       "How to create a will?",
		Description: "A step-by-step guide to creating a will and ensuring your assets are distributed according to your wishes.",
		Category:    "Estate Planning",
	}
	kb.AddArticle(article2)

	article3 := Article{
		Title:       "What is a trademark?",
		Description: "An explanation of what a trademark is and how it can protect your business.",
		Category:    "Intellectual Property",
	}
	kb.AddArticle(article3)

	// Get articles by category
	criminalLawArticles := kb.GetArticlesByCategory("Criminal Law")
	fmt.Println("Criminal Law Articles:")
	for _, article := range criminalLawArticles {
		fmt.Printf("- %s\n", article.Title)
	}

	// Search articles
	searchResults := kb.SearchArticles("create will")
	fmt.Println("Search Results:")
	for _, article := range searchResults {
		fmt.Printf("- %s\n", article.Title)
	}
	// // Search for articles
    // searchQuery := "Family Law"
    // searchResults := []Article{}
    // for _, article := range article {
    //     if strings.Contains(strings.ToLower(article.Title), strings.ToLower(searchQuery)) || strings.Contains(strings.ToLower(article.Description), strings.ToLower(searchQuery)) {
    //         searchResults = append(searchResults, article)
    //     }
    // }
    // fmt.Printf("\nSearch results for '%s':\n", searchQuery)
    // for _, article := range searchResults {
    //     fmt.Printf("\t%s: %s\n", article.Title, article.Description)
    // }

// This piece of code is an example of how to search for articles within a slice of Article values based on a search query.

// The first line defines the search query, in this case "Family Law".

// The second line initializes an empty slice to hold the search results.

// Then, the code loops through each Article in the articles slice. For each Article, it checks whether the search query string is contained within the article's Title or Description fields, using the strings.Contains() function. If the query string is found in either field, the Article is appended to the searchResults slice.

// Finally, the code prints out the search results, iterating over the searchResults slice and printing out the Title and Description fields of each matching Article.

// You can modify this code to fit your specific needs, for example by using a different search query string or by searching for the query in different fields of the Article struct.
}



// This code defines two functions: GetArticlesByCategory and SearchArticles. The former returns a list of all articles in the knowledge base that belong to a specific category, while the latter returns a list of all articles that contain a given search query in their title or description. You can use these functions to implement categorization and search functionalities in your content management system.