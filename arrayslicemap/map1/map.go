package main

import (
	"fmt"
)

func main() {
	//var aprovados map[int]string
	//map deve ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[12345679] = "Pedro"
	aprovados[12345677] = "Jose"

	//fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%-10s (CPF: %d)\n", nome, cpf)
	}

	mapTest := make(map[int]string)

	mapTest[1] = "A"

	fmt.Println(mapTest)

	fmt.Println(aprovados[12345678])
	delete(aprovados, 12345678)
	fmt.Println(aprovados[12345678])
	fmt.Println(aprovados[12345677])

}
