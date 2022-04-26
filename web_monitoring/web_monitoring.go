package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	for {
		line()
		menu()

		command := command()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 3:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando...")
			os.Exit(-1)
		}
	}
}

func menu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do Programa")
}

func command() int {
	var inputCommand int
	fmt.Scan(&inputCommand)
	line()
	fmt.Println("O comando escolhido foi: ", inputCommand)

	return inputCommand
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	line()
	fmt.Println(resp)
	line()

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func line() {
	fmt.Println("")
}
