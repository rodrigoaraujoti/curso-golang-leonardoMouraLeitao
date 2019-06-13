// AULA_065 - CRIANDO E INSTALANDO UM PACOTE DO GITHUB
// go get -u github.com/rodrigoaraujoti/go_area
// o comando acima ira baixar o pacote go_area do github para a pasta gopath (user/go)
// em seguida, o codigo abaixo ira utilizar este pacote
package main

import (
	"fmt"

	"github.com/rodrigoaraujoti/go_area/goarea"
)

func main() {
	fmt.Println(goarea.Circ(4.0))
}
