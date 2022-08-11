package main

import "fmt"

func main() {
	x := 1
	y := 2

	x++
	y--

	fmt.Println(x, y)

	//Não e permitido em go
	//fmt.Println(x == y--)
}
