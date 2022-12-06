package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Books)
}
