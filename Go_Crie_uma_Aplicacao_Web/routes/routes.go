package routes

import (
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Crie_uma_Aplicacao_Web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
