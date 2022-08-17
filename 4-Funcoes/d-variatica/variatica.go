package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	mediaAluno := media(10, 9, 5, 4, 8, 10)
	fmt.Printf("Media: %.2f", mediaAluno)
}
