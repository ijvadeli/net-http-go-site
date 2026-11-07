package main

import (
	"net/http"
)

func main() {
	// Create router instance so any third party packages can't create new routes
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Use the router
	print("Hosting on port :8080")
	http.ListenAndServe(":8080", router)
}
