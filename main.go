package main

import (
	"log"
	"net/http"
	"path/filepath"

	server "server/src"
)

// main functin where thr program will start executing
// the function will run the server and start lisening to requests
// from the browser a request come with method "GET" and root "/"
// the htttp.Handlefunc playyes a rule of guid => we give it  the root and check it with the req root if they math it send that req to specify func


func main() {
	http.HandleFunc("/", server.HomePage)
	http.HandleFunc("/details/", server.SecondPage)
	http.HandleFunc("/lastpage/", server.LastPage)
	staticDir := filepath.Join(".", "static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
