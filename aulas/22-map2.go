package main

import (
	"fmt"
	"strings"
)

func main () {
	// Criando um map para contar as palavras
	text := "go é uma linguagem de programação e go é muito rápida e é parecido com C++"
	words := strings.Fields(text)

	wordCount := make(map[string]int)

	// Contagem de frequencia de palavras
	for _, word := range words {
		wordCount[word]++
	}

	// Imprimir as frequencias
	fmt.Println("Contagem de palavras")
	for word, count := range wordCount {
		fmt.Printf("Palavra %s | Frequência %d \n", word, count)
	}
}