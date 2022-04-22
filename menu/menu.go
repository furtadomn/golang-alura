package main

import "fmt"

func main() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")

	var command int

	// fmt.Scanf("%d", &command)
	fmt.Scan(&command) // dessa forma não é necessário passar o modificador %d
	fmt.Println("O comando escolhido foi:", command)
	fmt.Println("O endereço da minha variável de comando é:", &command)
}
