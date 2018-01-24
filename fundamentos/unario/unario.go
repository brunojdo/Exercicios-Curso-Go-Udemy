package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2

	//apenas postfix (pós fixada)
	x++ // x = x +1
	fmt.Println(x)

	y-- // y = y -1
	fmt.Println(y)

	// isso não pode, usar dentro de uma expressão
	//fmt.Println(x == y--)

}
