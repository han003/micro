package main

import (
	"log"
	"net/http"
)

const (
	port = ":9090"
)

func main() {
	log.Println("Hello in main")

	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello world")
	})

	http.ListenAndServe(port, nil)
}
