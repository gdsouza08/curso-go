package main

import "fmt"

func main(){
	var notas = [5] float64 {8.7, 6.8, 9.3, 6.4, 7.9}

	soma := notas[0] + notas[1] + notas[2] + notas[3] + notas[4]

	media := soma / float64(len(notas))

	fmt.Println("Soma das notas é:", soma)
	fmt.Printf("Média das notas é: %.2f", media)
}