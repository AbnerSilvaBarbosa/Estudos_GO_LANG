package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonomos
	turboLigado bool
}

func main() {
	f := ferrari{
		turboLigado: true,
	}
	f.nome = "F40"
	f.velocidadeAtual = 50

	fmt.Println(f)
}
