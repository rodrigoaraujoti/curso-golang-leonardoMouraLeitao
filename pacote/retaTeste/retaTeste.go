// Testando a visibilidade de itens do pacote "reta"
package main

import "fmt"

func main() {
	p1 := reta.Ponto{2, 2}
	p2 := reta.Ponto{2, 4}

	// fmt.Println(catetos(p1, p2))
	fmt.Println(reta.Distancia(p1, p2))
}
