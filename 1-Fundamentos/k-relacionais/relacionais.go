package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 1 < 2)
	fmt.Println(">", 1 > 3)
	fmt.Println("<=", 1 <= 2)
	fmt.Println(">=", 1 >= 2)

	d1 := time.Unix(1, 6)
	d2 := time.Unix(1, 6)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome  string
		Idade int
	}

	p1 := Pessoa{"Abner", 16}
	fmt.Println(p1)
}
