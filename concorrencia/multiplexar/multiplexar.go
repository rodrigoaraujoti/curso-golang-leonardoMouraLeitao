// AULA76 - PADRAO DE CONCORRENCIA: MULTIPLEXADOR
// Juntar multiplos canais em um canal
package main

import (
	"fmt"

	"github.com/rodrigoaraujoti/go_area/htmlutil"
)

// encaminhar dado de um canal de origem para um canal de destino
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) em um canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		htmlutil.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		htmlutil.Titulo("https://g1.globo.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
