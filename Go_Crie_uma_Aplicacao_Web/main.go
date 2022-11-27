package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{
			Nome:       "Camiseta",
			Descricao:  "Azul bem bonita",
			Preco:      49.90,
			Quantidade: 10,
		},
		{"Camiseta", "Confortável", 250.00, 5},
		{"Fone", "Muito bom", 195.00, 8},
		{"Garrafa de água", "2 litros", 75.90, 5},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
