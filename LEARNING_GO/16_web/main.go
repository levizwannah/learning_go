package main

import (
	"fmt"
	"net/http"
)

func showIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func main() {
	fmt.Println("Starting Server...")
	http.HandleFunc("/", showIndex)
	http.ListenAndServe(":3000", nil)
}
