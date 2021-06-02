package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

// Create new record using Post method
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	//get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	fmt.Fprintf(w, "%+v", string(reqBody))

}

//Delete the data from record using Detele method
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	//once again we will need to parse the path parameters
	vars := mux.Vars(r)
	//we will need to extract the 'id' of the article we
	// wish to delete
	id := vars["id"]
	for index, article := range Articles {
		//if our id path parameter matched one of our
		//articles
		if article.Id == id {
			//updates our Articles array to remove the article
			Articles = append(Articles[:index], Articles[index+1:]...)
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

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8090", myRouter))
}

func main() {
	Articles = append(Articles,
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"})
	Articles = append(Articles, Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"})

	handleRequest()
}
