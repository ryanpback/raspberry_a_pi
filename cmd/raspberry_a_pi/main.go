package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world! This is Go!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Println("Listing for requests at http://localhost:9700")
	log.Fatal(http.ListenAndServe(":9700", nil))
}
