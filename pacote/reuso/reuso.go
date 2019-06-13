// AULA_064 - REUTILIZANDO PACOTE EXTERNO
// - No exemplo abaixo o pacote foi criado diretamente no diretorio go_path
// - posteriormente eu deletei este, uma vez que criei o mesmo codigo no github rodrigoaraujoti/go_area
package main

import (
	"fmt"

	"github.com/rodrigoaraujoti/go_area/goarea"
)

func main() {
	fmt.Println(goarea.Circ(6))
	fmt.Println(goarea.Rect(5, 2))
	// fmt.Println(area._TrianguloEq(5, 2)) esta privado
}
