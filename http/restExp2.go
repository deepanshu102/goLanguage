package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = make([]Article, 0)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	fmt.Println("Entpoint hit:Homepage")
}

// return all Articles

func returnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// Return a Single Article  according to ID as pathvariable
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(vars, key)
	fmt.Fprintf(w, "key: "+key)

	//Loop over all of our Articles
	//if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	//return all Articles
	myRouter.HandleFunc("/articles", returnAllArticles)
	// return Single Article
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

func main() {
	Articles = append(Articles,
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"})
	Articles = append(Articles, Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"})

	handleRequest()
}
