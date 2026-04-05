package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Endereco string
}
func main (){
	pessoa1 := Pessoa{
		Nome: "Maira Dias",
		Idade: 18,
		Endereco: "Rua do caminho 877",
	}
	fmt.Println("Informações do carro:")
	fmt.Printf("Nome: %s \n", pessoa1.Nome)
	fmt.Printf("Idade: %d \n", pessoa1.Idade)
	fmt.Printf("Endereço: %s \n", pessoa1.Endereco)
}
