// AULA_74 PADRAO DE CONCORRENCIA: GENERATOR
package main

import (
	"fmt"
	"reflect"

	"github.com/rodrigoaraujoti/go_area/htmlutil"
)

func main() {
	urls := []string{
		"https://www.cod3r.com.br",
		"https://www.google.com.br",
		"https://g1.globo.com/", //https://www.amazon.com ->erro sem tratamento
		"https://www.youtube.com",
	}

	t1 := htmlutil.Titulo(urls...)

	fmt.Println("--- ranking")
	for i := 0; i < cap(urls); i++ {
		fmt.Println(i, <-t1)
	}

	fmt.Println("\n--- testes sobre tipo: <-chan string")
	fmt.Println("TypeOf <-chan string", reflect.TypeOf(t1))
	fmt.Println("cap of <-chan string", cap(t1))
}
