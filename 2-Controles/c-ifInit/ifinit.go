package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i >= 5 { //Variavel e estanciada dentro do if e usada dentro do escopo
		fmt.Println("Nice")
	} else {
		fmt.Print("Não foi dessa vez")
	}
}
