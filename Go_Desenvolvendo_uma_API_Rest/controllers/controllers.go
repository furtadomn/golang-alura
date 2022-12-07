package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, book := range models.Books {
		if strconv.Itoa(book.ID) == id {
			json.NewEncoder(w).Encode(book)
		}
	}
}
