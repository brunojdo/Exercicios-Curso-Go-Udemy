package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	//convertendo um inteiro para flutuante
	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	//cuidado...
	//nesse caso vai pegar o caracter correspondente na tabela unicode
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("TEste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	//b, _ := strconv.ParseBool("1") = true...qualquer outro valor retorna falso

	if b {
		fmt.Println("Verdadeiro")
	}

}
