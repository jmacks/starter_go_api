package main

import (
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
