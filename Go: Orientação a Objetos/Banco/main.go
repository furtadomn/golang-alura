package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente."
	}
}

func main() {
	contaDaMarcella := ContaCorrente{}
	contaDaMarcella.titular = "Marcella"
	contaDaMarcella.saldo = 500

	fmt.Println(contaDaMarcella)
	fmt.Println(contaDaMarcella.Sacar(800))
	fmt.Println(contaDaMarcella.Sacar(400))
	fmt.Println(contaDaMarcella)
}
