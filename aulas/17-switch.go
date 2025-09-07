package main

import "fmt"

func main(){
	var nota float64

	fmt.Println("Digite a nota do aluno (0 a 10): ")
	fmt.Scan(&nota)

	switch {
	case nota >=7 && nota <=10:
		fmt.Println("Aprovado!")
	case nota >=5 && nota <7:
		fmt.Println("Recuperação!")
	case nota >=0 && nota <5:
		fmt.Println("Reprovado!")
	default: fmt.Println("Digite uma nota válida entre 0 e 10")
	}
}