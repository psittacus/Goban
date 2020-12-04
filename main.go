package main

import (
    "fmt"
    "log"
    "github.com/gorilla/mux"
    "net/http"
    "github.com/psittacus/Goban/data"
)

func HomeHandler( w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, data.GetAllTopics())
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
//    r.HandleFunc("/products", ProductsHandler)
//    r.HandleFunc("/articles", ArticlesHandler)
    log.Fatal(http.ListenAndServe(":8080", r))
}

