package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido:%d", n)
	case n == 0:
		return 1, nil

	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, error := fatorial(1)
	if error != nil {
		fmt.Print(error)
	} else {

		fmt.Print(resultado)
	}

}
