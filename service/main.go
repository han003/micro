package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":9090"
)

func main() {
	log.Println("Hello in main")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("HANDLE")

		fmt.Fprintf(rw, "Hello")
	})

	http.ListenAndServe(port, nil)
}
