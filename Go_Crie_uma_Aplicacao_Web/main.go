package main

import (
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Crie_uma_Aplicacao_Web/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
