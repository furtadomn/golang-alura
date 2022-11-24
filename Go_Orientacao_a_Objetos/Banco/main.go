package main

import (
	"fmt"

	"github.com/furtadomn/golang-studies/Go_Orientacao_a_Objetos/Banco/contas"
)

func main() {
	contaDaMarcella := contas.ContaCorrente{Titular: "Marcella", Saldo: 1300}
	contaDoChristian := contas.ContaCorrente{Titular: "Christian", Saldo: 900}

	status := contaDaMarcella.Transferir(200, &contaDoChristian)
	fmt.Println(status)

	fmt.Println(contaDaMarcella)
	fmt.Println(contaDoChristian)
}
