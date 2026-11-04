package main

import (
	"net/http"
)

func main() {
	homeDir := http.Dir("./static")
	http.Handle("/", http.FileServer(homeDir))

	http.ListenAndServe(":8080", nil)
}
