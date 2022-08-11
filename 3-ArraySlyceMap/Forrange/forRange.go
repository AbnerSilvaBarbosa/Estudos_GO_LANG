package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador que vai contar o tamanho do array

	for i, numero := range numeros {
		fmt.Println("Index:", i, "Key:", numero)

	}

	for _, numero := range numeros {
		fmt.Println("Key:", numero)

	}

	for i := range numeros {
		fmt.Println("Index:", i)

	}
}
