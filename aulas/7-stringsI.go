package main

import 	(	
	"fmt"
	"strings"
)

func main(){
	movie1 := "As Branquelas"
	movie2 := "As branquelas"
	// Go é case sensitive
	fmt.Println(movie1 == movie2)

	movieDescription := `
		É um filme de dois agentes do FBI 
		que se infiltram em uma festa branca.
	`
	fmt.Println(movieDescription)
	
	line := "-"
	fmt.Println(strings.Repeat(line, 50))

	// Verifica se uma palavra existe dentro de uma string
	fmt.Println(strings.Contains(movieDescription, "FBI"))
	fmt.Println(strings.Contains(movieDescription, "Branca"))
}