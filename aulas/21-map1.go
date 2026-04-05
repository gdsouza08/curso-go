package main

import "fmt"

func main () {
	// Criando um map com nome do aluno como chave e a nota como valor
	estudantes := map[string]float64 {
		"Ana": 			8.4,
		"Guilherme":	9.9,
		"Maeva": 		7.2,
		"Chico": 		4.1,
	}
	fmt.Println("Classificação dos Alunos:")
	for nome, nota := range estudantes {
		status := "Reprovado"
		if nota >= 6.0 {
			status = "Aprovado"
		}
		fmt.Printf("Aluno: %s | Nota: %.2f | Status: %s\n", nome, nota, status)
	}
}