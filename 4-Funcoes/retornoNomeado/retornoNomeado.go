package main

import "fmt"

func trocar(p1, p2 int) (primeiro, segundo int) {
	primeiro = p2
	segundo = p1
	return //retorno limpo
}

func main() {
	resultadoTroca1, resultadoTroca2 := trocar(2, 3)

	fmt.Println(resultadoTroca1, resultadoTroca2)
}
