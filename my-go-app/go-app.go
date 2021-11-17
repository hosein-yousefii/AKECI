package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi, how is it going</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Let's see if it's ok!!!!</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":5555", nil)
}
