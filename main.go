package main

import (
	"ASCIIWEB"
	"fmt"
	"net/http"
)

func main() {
	// ()
	http.HandleFunc("/", ASCIIWEB.HomeHandler)
	http.HandleFunc("/ascii-art", ASCIIWEB.AsciiArtHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
		// ReadTimeout:  1000,
		// WriteTimeout: 1000,
	}
	fmt.Println("Starting server on port 8080")
	server.ListenAndServe()
}
