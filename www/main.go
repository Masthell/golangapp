package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// func handleRequest() {
//	http.HandleFunc("/", indexHandle)
//	http.ListenAndServe(":8080", nil)
// }

func main() {
	// handleRequest()
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    mux := http.NewServeMux()

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    mux.HandleFunc("/", indexHandle)
    http.ListenAndServe(":"+port, mux)
}

// localhost:8080