package main

import "fmt"

func resultadoFinal(n float64) string {

	nota := n
	switch nota {
	case 1:
		return "E"

	case 2:
		return "F"
	}

	return "Sem resultados"

}

func main() {
	result := resultadoFinal(1)
	fmt.Println(result)
}
