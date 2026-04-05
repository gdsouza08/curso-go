package main

import "fmt"

func main (){
	notas := []float64{8.5, 9.0, 7.5, 6.8, 9.3}

	fmt.Println("Cálculo da média com slice")

	var total float64
	for _, nota := range notas {
		total += nota // total = total + notas
	}

	media := total / float64(len(notas))
	fmt.Println("Soma das notas: ", total)
	fmt.Println("Média das notas: ", media)
}