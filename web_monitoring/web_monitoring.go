package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for {
		line()

		sitesReader()

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
	line()

	return inputCommand
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento...")
	line()

	sites := sitesReader()

	for i, site := range sites {
		fmt.Println("Posição: ", i, "Site: ", site)
		siteTest(site)
	}

}

func siteTest(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
	line()
}

func sitesReader() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()

	return sites
}

func line() {
	fmt.Println("")
}
