package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/Users/jiekexuan/go/learn.go")))
	log.Fatal(http.ListenAndServe(":8089", nil))
}
