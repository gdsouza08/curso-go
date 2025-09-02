package main

import (
	"fmt"
	"strings"
)

func main(){
	movie1 := "As Branquelas"
	movie2 := "As branquelas"
	fmt.Println(movie1 == movie2)

	movieDescription := `
		É um filme de dois agentes do FBI 
		que se infiltram em uma festa branca.
	`

	// Convertendo texto para maiúsculo
	fmt.Println(strings.ToUpper(movieDescription))
	// Convertendo texto para minúsculo
	fmt.Println(strings.ToLower(movieDescription))
	// Primeira letra em maiúsculo
	fmt.Println(strings.Title(movieDescription))

	// Encontrar a posição de um caractere
	fmt.Println(strings.Index(movieDescription, "b"))

	// Contando o número de ocorrências de um caractere
	fmt.Println(strings.Count(movieDescription, "a"))
	fmt.Println(strings.Count(movieDescription, "b"))

	// Substitiu um elemento por outro
	fmt.Println(strings.ReplaceAll(movieDescription, "filme", "série"))
}