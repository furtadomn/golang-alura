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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso! Saldo em conta:", c.saldo
	} else {
		return "O valor do depósito não pode ser negativo! Saldo em conta:", c.saldo
	}
}

func main() {
	contaDaMarcella := ContaCorrente{}
	contaDaMarcella.titular = "Marcella"
	contaDaMarcella.saldo = 500

	fmt.Println(contaDaMarcella)
	fmt.Println(contaDaMarcella.Sacar(400))
	fmt.Println(contaDaMarcella)

	status, valor := contaDaMarcella.Depositar(500)
	fmt.Println(status, valor)
}
