package main

import (
	"fmt"
)

type esportivo interface {
	ligarturbo()
}

type ferrari struct {
	modelo          string
	turboligado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarturbo() {
	f.turboligado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarturbo()

	// quando trabalhar diretamente
	//com interface tem que passar como referencia
	//utilizando o &
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarturbo()

	fmt.Println(carro1, carro2)

}
