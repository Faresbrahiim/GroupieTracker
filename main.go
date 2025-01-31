package main

import (
	"log"
	"net/http"
	"path/filepath"

	server "server/src"
)

func main() {
	http.HandleFunc("/", server.HomePage)
	http.HandleFunc("/details/", server.SecondPage)
	http.HandleFunc("/lastpage/", server.LastPage)
	staticDir := filepath.Join(".", "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
