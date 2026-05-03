package main 

import "fmt"

var nameClient string

// 1 - Imprimir mensagem de boas vindas
func welcome () {
	fmt.Println("Digite seu nome: ")
	fmt.Scanln(&nameClient)
	fmt.Println("Bem vindo ao sistema de filmes " + nameClient + "!!!")
}

func createMovie () {
	var name string
	var yearRelease int
	var moviePrice float64
	var dummy string

	fmt.Println("Digite o nome do filme: ")
	fmt.Scanln(&name)
	// Limpar o buffer de entrada
	fmt.Scanln(&dummy)

	fmt.Println("Digite o ano de lançamento do filme: ")
	fmt.Scanln(&yearRelease)

	fmt.Println("Digite o preço do filme: ")
	fmt.Scan(&moviePrice)

	fmt.Printf("%s (%d) - R$ %.2f\n", name, yearRelease, moviePrice)
}

func main () {
	welcome()
	fmt.Println("**Utilizando Função**")
	createMovie()
}