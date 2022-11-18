package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/home", home_handler)
	http.HandleFunc("/profile", profile_handler)
	http.ListenAndServe(":8080", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello newbees</h1>")
	fmt.Fprintln(w, "<h2>hello newbees</h2>")
	fmt.Fprintln(w, "<h3>hello newbees</h3>")
}
func home_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello newbees this is home ")
}
func profile_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello newbees this is profile")
}
