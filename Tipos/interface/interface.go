package main

import (
	"fmt"
	"strconv"
)

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco int
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return p.nome + " " + strconv.Itoa(p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa = pessoa{"Abner", "Silva"}
	imprimir(coisa)
}
