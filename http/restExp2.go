package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = make([]Article, 0)

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	fmt.Println("Entpoint hit:Homepage")
}
func handleRequest() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
func main() {
	Articles = append(Articles,
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"})
	Articles = append(Articles, Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"})

	handleRequest()
}
