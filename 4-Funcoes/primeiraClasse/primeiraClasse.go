package main

import "fmt"

func main() {
	soma := func(a, b int) int {
		return a + b
	}

	fmt.Println(soma(1, 2))

}
