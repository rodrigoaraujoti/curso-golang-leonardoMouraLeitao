// AULA_061 - TIPO INTERFACE
package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} // usando interface anonima como um tipo generico
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{} // criando um tipo que eh do tipo interface vazia, tambem serve como um tipo generico

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do Google"}
	fmt.Println(coisa2)

}
