package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	greeting()

	commandsList()
}

func greeting() {
	var name string

	line()
	fmt.Println("   Olá! Qual o seu nome?")
	fmt.Print("   ~> ")
	fmt.Scan(&name)

	clear()

	line()
	fmt.Println("   BEM VINDE À LINGUAGEM GO,", strings.ToUpper(name), "!")
	line()
}

func commandsList() {
	fmt.Println("   Digite um dos comandos abaixo:")
	line()
	fmt.Println("   [1] O que é Go?")
	fmt.Println("   [2] Consultar a documentação do Go")
	fmt.Println("   [3] Quero aprender Go!! ")
	fmt.Println("   [4] Sair do Programa")

	commands()
}

func commands() {
	var command int

	line()
	fmt.Print("   ~> ")
	fmt.Scan(&command)

	clear()

	switch command {
	case 1:
		line()
		goDescription()
		line()
		exitOrBackToMenuList()
	case 2:
		goDocs()
		exitOrBackToMenuList()
	case 3:
		goCourses()
		exitOrBackToMenuList()
	case 4:
		exit()
	default:
		invalidCommand()
		exitOrBackToMenuList()
	}
}

func exitOrBackToMenuList() {
	line()
	fmt.Println("   [1] Voltar")
	fmt.Println("   [2] Sair do Programa")
	fmt.Print("   ~> ")
	exitOrBackToMenuCommands()
}

func exitOrBackToMenuCommands() {
	var menucommands int

	fmt.Scan(&menucommands)
	clear()

	switch menucommands {
	case 1:
		commandsList()
	case 2:
		exit()
	default:
		invalidCommand()
		exitOrBackToMenuList()
	}
}

func goDescription() {
	line()
	fmt.Println("   ~~ LINGUAGEM GO ~~")
	line()
	fmt.Println("   Criado pela Google em 2009, o Go é uma linguagem de programação")
	fmt.Println("   compilada, simultânea, com coleta de lixo e tipada estaticamente.")
	line()
}

func goCourses() {
	line()
	fmt.Println("   ~~ CURSOS ~~")
	line()
	fmt.Println("   - Alura: https://www.alura.com.br/cursos-online-programacao/golang")
	fmt.Println("   - Udemy: https://www.udemy.com/course/curso-go/")
	fmt.Println("   - Youtube (*GRÁTIS*): https://www.youtube.com/c/AprendaGo/videos")
	line()
	fmt.Println("   * Utilize o comando ctrl + click para acessar os links acima *")
	line()
}

func goDocs() {
	line()
	fmt.Println("   ~~ DOCUMENTAÇÃO ~~")
	line()
	fmt.Println("   - Documentação oficial: https://go.dev/doc/ (ctrl + click)")
	line()
}

func exit() {
	line()
	fmt.Println("   Saindo... Até logo e bons estudos!")
	line()
}

func invalidCommand() {
	line()
	fmt.Println("   ** Favor digitar um comando válido **")
	line()
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func line() {
	fmt.Println("")
}
