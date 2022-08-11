package main

import "fmt"

//não tem operador ternario

func obeterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}

}

func main() {
	resultado := obeterResultado(7.5)
	fmt.Println(resultado)
}
