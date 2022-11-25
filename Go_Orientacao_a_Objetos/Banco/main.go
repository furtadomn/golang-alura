package main

import (
	"fmt"

	"github.com/furtadomn/golang-studies/Go_Orientacao_a_Objetos/Banco/contas"
)

func main() {
	contaDaMarcella := contas.ContaPoupanca{}
	contaDaMarcella.Depositar(100)
	PagarBoleto(&contaDaMarcella, 60)

	fmt.Println(contaDaMarcella.ObterSaldo())

	contaDoChris := contas.ContaCorrente{}
	contaDoChris.Depositar(500)
	PagarBoleto(&contaDoChris, 600)

	fmt.Println(contaDoChris.ObterSaldo())
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
