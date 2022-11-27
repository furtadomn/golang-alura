package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
