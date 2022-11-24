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

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaMarcella := ContaCorrente{titular: "Marcella", saldo: 1300}
	contaDoChristian := ContaCorrente{titular: "Christian", saldo: 900}

	status := contaDaMarcella.Transferir(200, &contaDoChristian)
	fmt.Println(status)

	fmt.Println(contaDaMarcella)
	fmt.Println(contaDoChristian)
}
