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
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
    articles := Articles{
        Article{Title:"Test Title", Desc: "Test Description", Content: "Hello World"},
    }

    fmt.Println("Endpoint Hit: All articles Endpoint")
    json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Test POST endpoint worker")
}

func handleRequests() {

    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", allArticles).Methods("GET")
    myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
    handleRequests()
}