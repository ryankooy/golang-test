package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Dock Yourself")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hell")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
