package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Abner", "Abner", "Abner", "Abner", "Abner", "Abner", "Abner", "Abner", "Abner"}
	aprovados = append(aprovados, "Rafael")
	imprimirAprovados(aprovados...)
}
