package main

import (
	"fmt"
)

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
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10}, //passar informando os nomes do campo
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	//pedido.itens = append(pedido.itens, item{1, 1, 12.0})
	//pedido.itens = append(pedido.itens, item{1, 1, 18.0})

	fmt.Printf("VAlor total do pedido é R$%.2f", pedido.valorTotal())

	//fmt.Println("\n", len(pedido.itens), cap(pedido.itens))
}
