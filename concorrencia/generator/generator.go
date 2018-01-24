package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - Canal somente leitura. Quando o canal é
// informado como "<-chan", siginifica que é um
//canal que não vai receber dados. Vamos apenas
//consumir os dados deste canal
func titulo(ursls ...string) <-chan string {
	c := make(chan string)
	for _, url := range ursls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.intel.com", "https://www.globo.com")
	t2 := titulo("https://www.cod3r.com", "https://www.ibm.com")
	fmt.Println("Primeiros", <-t1, "|", <-t2)
	fmt.Println("Segundos", <-t1, "|", <-t2)

}
