package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = m.Pi
	var raio = 3.2

	//formsa reduzida de criar uma var
	//Declarando e atribuindo a variavel
	area := PI * m.Pow(raio, 2)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	var temUm, temDois bool = true, false

	g, h, i := 2, false, "epa!"

	fmt.Println("A área da circunferência é:", area, c, d)
	fmt.Println(temUm, temDois)
	fmt.Println(g, h, i)
}
