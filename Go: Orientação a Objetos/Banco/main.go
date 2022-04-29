package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaMarcella := ContaCorrente{"Marcella", 123, 56789, 125.70}
	fmt.Println(contaDaMarcella)

	var contaDoChris *ContaCorrente
	contaDoChris = new(ContaCorrente)
	contaDoChris.titular = "Christian"

	fmt.Println(*contaDoChris)
}
