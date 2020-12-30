package main

import (
    "log"
    "github.com/gorilla/mux"
    "net/http"
    "github.com/psittacus/Goban/data"
    "fmt"
)

func HomeHandler( w http.ResponseWriter, r *http.Request) {
    resp, err := data.GetAllTopics()
    if err != nil {
        log.Fatal(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w,resp)
    return
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
//    r.HandleFunc("/products", ProductsHandler)
//    r.HandleFunc("/articles", ArticlesHandler)
    log.Fatal(http.ListenAndServe(":8080", r))
}

