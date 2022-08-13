package main

import (
	"fmt"
	"time"
)

func falar(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Println(pessoa + ": " + texto)
	}
}

func main() {
	go falar("Abner", "Gosto de sorvete", 5)
	falar("Marcos", "Gosto de pudim", 5)
}
