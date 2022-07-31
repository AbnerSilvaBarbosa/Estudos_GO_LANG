package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompreto string) {
	parte := strings.Split(nomeCompreto, " ")
	p.nome = parte[0]
	p.sobrenome = parte[1]
}

func main() {
	p1 := pessoa{
		nome:      "Abner",
		sobrenome: "Silva",
	}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Marcos Souza")
	fmt.Println(p1.getNomeCompleto())
}
