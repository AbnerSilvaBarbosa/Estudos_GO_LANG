package main

import "fmt"

// (nome tipo) tipo do returno
func somar(a int, b int) int {

	return a + b

}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	imprimir(somar(10, 20))
}
