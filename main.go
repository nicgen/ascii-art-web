package main

import (
	"ASCII"
	"ASCIIWEB"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", ASCIIWEB.HomeHandler)
	http.HandleFunc("/ascii-art", ASCIIWEB.AsciiArtHandler)
	http.HandleFunc("/download", ASCII.DownloadHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	server := http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadHeaderTimeout: 10 * time.Second,  // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second,  // temps maximum d'écriture de la réponse
		IdleTimeout:       120 * time.Second, // temps maximum entre deux rêquetes
		MaxHeaderBytes:    1 << 20,           // 1 MB // maxinmum de bytes que le serveur va lire
		ReadTimeout:       1000,
		// WriteTimeout:      1000,
	}
	fmt.Println("Starting server on port 8080")
	// server.ListenAndServe()
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
