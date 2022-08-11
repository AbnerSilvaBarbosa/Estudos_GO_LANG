package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv20 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv20, comprarSorvete

}

func main() {
	tv50, tv20, sorvete := compras(true, false)
	fmt.Println("Tv-50:", tv50)
	fmt.Println("Tv-20:", tv20)
	fmt.Println("Sorvete:", sorvete)

}
