package main

import "fmt"

func imprimirResultado(nota float64, aluno bool) string {
	if nota >= 5 && aluno {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	resultado := imprimirResultado(5, true)

	fmt.Println(resultado)
}
