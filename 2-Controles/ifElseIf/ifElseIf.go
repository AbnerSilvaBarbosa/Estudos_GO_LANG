package main

import "fmt"

func notaParaConceito(nota float64) string {

	if nota == 10 || nota >= 9 {
		return "A"
	} else if nota >= 7 || nota <= 9 {
		return "B"
	} else if nota >= 6 || nota <= 5 {
		return "C"
	} else if nota >= 4 || nota <= 3 {
		return "D"
	} else if nota >= 2 || nota <= 1 {
		return "E"
	} else {
		return "F"
	}

}

func main() {
	resultado := notaParaConceito(10)
	fmt.Println(resultado)
}
