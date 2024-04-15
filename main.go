package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

var  Articles[]Article

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our first go API")
	fmt.Println("Endpoint hit: home")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// This route handler is writing natively
// func handleRequests() {
// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/articles", returnAllArticles)
// 	log.Fatal(http.ListenAndServe(":7000", nil))
// }

// func main() {
// 	Articles = []Article{
// 		Article{Title:"Hello World", Desc: "Say hello to the world", Content: "kuda"},
// 		Article{Title:"Hello World", Desc: "Say hello to the world", Content: "kuda"},
// 	}
// 	handleRequests()
// }

// Here are the example of using mux router
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/all", returnAllArticles)

	log.Fatal(http.ListenAndServe(":7000", router))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title:"Hello World", Desc: "Say hello to the world", Content: "kuda"},
		Article{Title:"Hello World", Desc: "Say hello to the world", Content: "kuda"},
	}
	handleRequests()
}
