package main

import (
	"fmt"
	"time"

	"github.com/owilliammartins/paginaWeb"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := paginaWeb.Titulo(url1)
	c2 := paginaWeb.Titulo(url2)
	c3 := paginaWeb.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(4000 * time.Millisecond):
		return "Demorou mais de 1 segundo."
		/* 	default:
		return "Sem retorno" */
	}
}

func main() {
	campeao := oMaisRapido("https://GOOGLE.com", "https://amazon.com", "https://yahoo.com")
	fmt.Println(campeao)
}
