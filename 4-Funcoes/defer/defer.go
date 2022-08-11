package main

import "fmt"

func obeterValorAprovado(num int) int {
	defer fmt.Println("Fim!")

	if num >= 5000 {
		defer fmt.Println("Valor alto....")
		return 5000
	} else {
		fmt.Println("Valor baixo")
		return num
	}
}

func main() {
	fmt.Println(obeterValorAprovado(3000))
}
