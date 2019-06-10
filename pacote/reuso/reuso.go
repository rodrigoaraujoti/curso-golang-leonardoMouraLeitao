// AULA_064 - REUTILIZANDO PACOTE EXTERNO
// - No exemplo abaixo o pacote foi criado diretamente no diretorio go_path
// - posteriormente eu deletei este, uma vez que criei o mesmo codigo no github rodrigoaraujoti/go_area
package main

import (
	"fmt"

	"github.com/rodrigoaraujoti/go_area"
)

func main() {
	fmt.Println(go_area.Circ(6))
	fmt.Println(go_area.Rect(5, 2))
	// fmt.Println(area._TrianguloEq(5, 2)) esta privado
}
