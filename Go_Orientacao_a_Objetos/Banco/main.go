package main

import (
	"fmt"

	"github.com/furtadomn/golang-studies/Go_Orientacao_a_Objetos/Banco/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
