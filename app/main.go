package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func PostModel(js string) (string, int) {
	err := 0
	params, err := ParseJSON(js)
	if err != 0 {
		return "", 1
	}
	randStr, err := Generate(&params)
	if err != 0 {
		return "", 1
	}
	id, err := PutRand(randStr, db)

	return id, err
}

func GetModel(id string) (string, int) {
	randVal, err := GetRand(id, db)
	return randVal.val, err
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var err error
	port := os.Args[1]
	db, err = sql.Open("sqlite3", DBPATH)
	http.HandleFunc("/api/generate/", postHandler)
	http.HandleFunc("/api/retrieve/", getHandler)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
