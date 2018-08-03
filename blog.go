package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Post is a blog post struct
type Post struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Posts is an array all post structures
type Posts []Post

// AllPosts retrieves all blog posts from db &
//  encodes them to JSON format
func AllPosts(w http.ResponseWriter, r *http.Request) {
	posts := Posts{
		Post{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(posts)
}

// TestPosts is a placeholder for tests
func TestPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}
