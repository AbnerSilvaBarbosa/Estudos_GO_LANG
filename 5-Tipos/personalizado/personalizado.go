package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	}
	return "Teste"
}

func main() {

	var nota nota = 10
	fmt.Println(notaParaConceito(nota))
}
