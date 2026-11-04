package main

import (
	"fmt"
	"html"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path: %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/", testHandler)
	http.ListenAndServe(":8080", nil)
}
