package main

import (
	f "fmt"
	m "math"
)

func main() {

	x := m.Pi
	xs := f.Sprint(x)

	f.Println("O valor de x é ", x)

	//Só e possivel concatenar com a mesma tipagem
	f.Println("O valor de x é " + xs)

	//2 casas decimais
	f.Printf("O valor de x é %.2f", x)
}
