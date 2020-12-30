package data

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    _ "fmt"
    "log"
)

const (
    sqliteDbPath = "./data/board.sqlite3"
)

type Source interface {
    //GET Request to the root "/"
    GetAllTopics() (string, error)
}

func GetAllTopics() (string, error) {
   resp := ""

   db, err := sql.Open("sqlite3", sqliteDbPath)
   if err != nil {
        log.Fatal(err)
        return "", err
   }
   defer db.Close()

    //Get all database names
    rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")

   if err != nil {
        log.Fatal(err)
        return "", err
   }
    defer rows.Close()
    for rows.Next() {
        var temp string
        if err := rows.Scan(&temp); err != nil {
            log.Fatal(err)
        }
        resp+=temp+" "
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }

   return resp, err
}
