package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/furtadomn/golang-studies/Go_Crie_uma_Aplicacao_Web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		floatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		intQuantity, err := strconv.ParseInt(quantity, 10, 64)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CreateNewProduct(name, description, floatPrice, intQuantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.EditProduct(id)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Println("Erro na conversão do id:", err)
		}

		floatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		intQuantity, err := strconv.ParseInt(quantity, 10, 64)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.UpdateProduct(intID, intQuantity, name, description, floatPrice)
	}

	http.Redirect(w, r, "/", 301)
}
