package main

import "fmt"

// func main() {
// 	greeting()
// }

func main() {
	var first_name string = "Marcella "
	var last_name = "Furtado" // O Go consegue inferir o tipo das variáveis, então nós podemos omiti-lo
	age := 28                 // Declaração curta de variáveis

	fmt.Print("Olá! Meu nome é ", first_name, last_name)
	fmt.Print(". Eu tenho ", age, " anos.")
}
