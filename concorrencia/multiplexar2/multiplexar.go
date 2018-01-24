package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s está falando %d", pessoa, i)
		}
	}() // como a função é anonima, utilizando os parenteses já
	// esta fazendo a chamada da mesma. Outro modo de fazer
	// seria atribuir a função à uma variável e fazer a
	// chamada da mesma

	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := juntar(falar("João"), falar("maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)

}
