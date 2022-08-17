package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println(a + b)
	fmt.Println(a - b)

	fmt.Println(a / b)
	fmt.Println(a * b)

	fmt.Println(a % b)

	//valor binario
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)

	fmt.Println(math.Max(float64(a), float64(b)))
	fmt.Println(math.Min(float64(a), float64(b)))
	fmt.Println(math.Pow(float64(a), float64(b)))

}
