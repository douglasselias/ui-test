package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fs.ServeHTTP(response, request)
	})

	port := ":3000"
	log.Printf("Server started on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
