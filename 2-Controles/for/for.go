package main

import "fmt"

//Não existe while em Go

func main() {

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for b := 0; b < 5; b++ {
		fmt.Println("LALALA")
	}

}
