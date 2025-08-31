package main

import "fmt"

func main(){
	// Tipo de dado inteiro
	var idade int = 30
	// Tipo ponto flutuante
	var altura float64 = 1.75
	// Tipo de dado booleano (true or false)
	var maiorDeIdade bool = idade >=18
	// Tipo de dado string
	var nome string = "Guilherme"
	
	fmt.Println("Dados Pessoais")
	fmt.Println("Nome:")
	fmt.Println(nome)
	fmt.Println("Idade:")
	fmt.Println(idade)
	fmt.Println("Altura:")
	fmt.Println(altura)
	fmt.Println("maior de idade?")
	fmt.Println(maiorDeIdade)

	// Verifica tipo do dado
	fmt.Println(fmt.Sprintf("%T", maiorDeIdade))
}