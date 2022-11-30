package controllers

import (
	"html/template"
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Crie_uma_Aplicacao_Web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
