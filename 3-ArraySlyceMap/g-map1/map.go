package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345567] = "Abner"
	aprovados[12321221] = "Teste"
	fmt.Println(aprovados)

	for codigo, nome := range aprovados {
		fmt.Println(codigo)
		fmt.Println(nome)
	}

	delete(aprovados, 12321221)
	fmt.Println(aprovados)

}
