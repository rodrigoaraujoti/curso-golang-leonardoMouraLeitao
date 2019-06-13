// AULA_74 PADRAO DE CONCORRENCIA: GENERATOR
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Pattern - Rob Pike
// funcao que encapsula a goroutine

// <-chan - canal somente-leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) { //funcao anonima
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			// fmt.Println(reflect.TypeOf(r)) ~ obs: o vscode nao autocompleta este tipo: *regexp.Regexp
			c <- r.FindStringSubmatch(string(html))[1] //--> [0] toda a regex, [1] soh o coringa: (.*?)

		}(url) // executando funcao anonima
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com.br")
	t2 := titulo("https://g1.globo.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
