package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	//retorna a quantidade de núcleos fisicos
	//tem o processador
	a := runtime.NumCPU()

	fmt.Println(a)

	fmt.Println(reflect.TypeOf(a))
}
