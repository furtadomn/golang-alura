package main

import (
	"fmt"

	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/models"
	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/routes"
)

func main() {
	models.Books = []models.Book{
		{Title: "Se quiser mudar o mundo: Um guia político para quem se importa", Author: "Sabrina Fernandes", Pages: 192},
		{Title: "Terra das mulheres", Author: "Charlotte Perkins Gilman", Pages: 256},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
