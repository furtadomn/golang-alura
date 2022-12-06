package main

import (
	"fmt"

	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
