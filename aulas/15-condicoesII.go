package main

import "fmt"

func main(){
	var nota float32

	fmt.Println("Digite sua nota (0 a 10):")
	fmt.Scan(&nota)

	if nota <0 || nota >10{
		fmt.Println("Nota inválida! Digite uma nota entre 0 e 10")
	} else if nota >=7 {
		fmt.Println("Aprovado!")
	} else if nota >=5 {
		fmt.Println("Recuperação!")
	} else {
		fmt.Println("Reprovado!")
	}
}