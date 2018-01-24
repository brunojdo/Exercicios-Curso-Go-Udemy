package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/boom", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(500)
		res.Write([]byte("Boom!"))
	})
	go http.ListenAndServe(":9000", nil)

	//res, err := http.Get("http://127.0.0.1:9000/boom")
	res, err := http.Get("http://www.drogariaminasbrasil.com.br/")
	html, _ := ioutil.ReadAll(res.Body)
	r, _ := regexp.Compile("<title>(.*?)<\\/title>")
	titulo := r.FindStringSubmatch(string(html))[1]
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res.StatusCode: %d\n", res.StatusCode)
	fmt.Println(titulo)
}
