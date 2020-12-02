package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/view.html", "templates/edit.html", "templates/books.html"))

func main() {

	http.HandleFunc("/books/", bookHandler)
	http.HandleFunc("/hypermaze/", mazeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mazeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "edit")
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "books")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
