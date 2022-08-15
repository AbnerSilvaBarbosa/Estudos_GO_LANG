package main

import (
	"fmt"

	"github.com/AbnerSilvaBarbosa/area/area2"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}

}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		area2.Titulo("https://www.Google.com", "https://www.Youtube.com"),
		area2.Titulo("https://www.Google.com", "https://www.Youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
