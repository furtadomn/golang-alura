package routes

import (
	"log"
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/livros", controllers.AllBooks).Methods("GET")
	r.HandleFunc("/livros/{id}", controllers.GetBookByID).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
