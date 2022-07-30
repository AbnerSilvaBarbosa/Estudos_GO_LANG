package main

import "fmt"

func main() {
	// homogênea (mesmo tipo) e estática (fixo)

	var notas [3]float64

	notas[0], notas[1], notas[2] = 8.3, 9.5, 10

	conta1 := 0.0

	for i := 0; i < len(notas); i++ {
		conta1 += notas[i]
	}

	resultado := conta1 / float64(len(notas))

	fmt.Printf("Media: %.2f", resultado)

}
