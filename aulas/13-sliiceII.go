package main

import "fmt"

func main(){
	frutas := [] string {"Banana", "Amora", "Uva", "Manga", "Pêssego"}

	// Criando uma subslice pegando do indice de 1 a 3
	subslice := frutas [1:4]

	fmt.Println("Subslice de frutas:", subslice)

	subslice [0] = "Abacate"

	fmt.Println("Subslice:", subslice)

	fmt.Println("Slice após a modificação:", frutas)

}