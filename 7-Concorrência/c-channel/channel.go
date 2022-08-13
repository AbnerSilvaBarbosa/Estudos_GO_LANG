package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1   // enviando dados para o canal (escrita)
	a := <-ch // recebendo dados do canal (leitura)

	ch <- 2
	b := <-ch
	fmt.Println(a, b)

}
