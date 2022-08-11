package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {

	item1 := item{
		produtoID: 1,
		qtde:      2,
		preco:     10,
	}

	item2 := item{
		produtoID: 2,
		qtde:      5,
		preco:     2,
	}

	pedido := pedido{
		userID: 23,
		itens: []item{
			item1,
			item2,
		},
	}

	fmt.Print(pedido.valorTotal())

}
