package main

import (
	"fmt"
	"strconv"
	"strings"
)

type produto struct {
	codigo   int
	nome     string
	preco    float64
	desconto float64
	listEan  []string
}

var todosProdutos [2]produto

//Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func (p produto) codigoEAN(cont int) []string {
	listEan := make([]string, cont, cont*2)
	for i := 0; i < 3; i++ {
		listEan[i] = strconv.Itoa(i)
	}
	return listEan
}

func (p produto) novoCodigoEAN(listaEAN string) []string {
	return strings.Split(listaEAN, ",")
}

func main() {
	var produto1 produto
	produto1 = produto{
		codigo:   1,
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1)

	produto2 := produto{2, "Notebook", 2000, 0.10, nil}
	produto1.listEan = produto1.codigoEAN(25)

	produto2.listEan = produto2.novoCodigoEAN("12345,321654,142536")

	desconto := produto2.precoComDesconto()
	fmt.Println(produto2, desconto)

	todosProdutos[0] = produto1
	todosProdutos[1] = produto2

	fmt.Println(todosProdutos)

}
