package main

import (
	"fmt"
	"math/rand"
)

func main () {
	// gerador de números aletórios
	target := rand.Intn(10) + 1 // números entre 1 e 10
	fmt.Println("Jogo da Advinhação")
	fmt.Println("Tente adivinhar um número entre 1 e 10")

	var guess int

	for {
		fmt.Println("Digite seu palpite (ou 0 para sair)")
		fmt.Scan(&guess)

		if guess == 0 {
			fmt.Println("Saindo do jogo. o número era: ", target)
			break
		}

		if guess > target {
			fmt.Println("seu palpite é maior. Tente novamente!")
		} else if guess < target {
			fmt.Println("Seu palpite é menor. Tente novamente!")
		} else {
			fmt.Println("Parabéns!!! Você Acertou!")
		}
	}
}
