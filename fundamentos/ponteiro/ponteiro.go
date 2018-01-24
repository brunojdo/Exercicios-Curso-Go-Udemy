package main

import (
	"fmt"
)

func main() {
	i := 1

	// Go nao tem aritmetica de ponteiros
	var p *int = nil

	p = &i // pegando o endereco da variavel
	*p++   // desreferenciando (pegando o valor)

	i++

	//p++ isso não pode

	fmt.Println(p, *p, i, &i)

}
