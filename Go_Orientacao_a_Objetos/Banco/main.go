package main

import (
	"fmt"

	"github.com/furtadomn/golang-studies/Go_Orientacao_a_Objetos/Banco/clientes"
	"github.com/furtadomn/golang-studies/Go_Orientacao_a_Objetos/Banco/contas"
)

func main() {
	contaDaMarcella := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Marcella",
		CPF:       "123.123.123-13",
		Profissao: "Desenvolvedora",
	},
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         1000,
	}

	// clienteMarcella := clientes.Titular{"Marcella", "123.123.123-13", "Desenvolvedora"}
	// contaDaMarcella := contas.ContaCorrente{clienteMarcella, 123, 123456, 1000}

	fmt.Println(contaDaMarcella)
}
