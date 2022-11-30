package main

import (
	"html/template"
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Crie_uma_Aplicacao_Web/models"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
