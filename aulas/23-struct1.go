package main

import "fmt"

type Carro struct {
	Modelo string
	Ano int
	Cor string
}
func main (){
	carro1 := Carro{
		Modelo: "Fusca",
		Ano: 1958,
		Cor: "Vinho",
	}
	fmt.Println("Informações do carro:")
	fmt.Printf("Modelo: %s \n", carro1.Modelo)
	fmt.Printf("Ano: %d \n", carro1.Ano)
	fmt.Printf("Cor: %s \n", carro1.Cor)
}
