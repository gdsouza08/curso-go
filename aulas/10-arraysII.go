package main

import "fmt"

func main(){
	var frutas = [3] string {"Manga", "Abacaxi", "Morango"}

	fmt.Println("Array de frutas: ", frutas)
	fmt.Println("Array de frutas: ", frutas[0])
	fmt.Println("Array de frutas: ", frutas[1])
	fmt.Println("Array de frutas: ", frutas[2])
	fmt.Println("Array de frutas: ", frutas[0:2])
}