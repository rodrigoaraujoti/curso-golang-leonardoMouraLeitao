// AULA_020 - OPERADORES UNARIOS
package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas postfix
	x++ // operador unario de incremento postfix
	fmt.Println(x)

	y-- // operador unario de decremento postfix
	fmt.Println(y)

	// fmt.Println(x == y--) --> operadores unarios nao sao permitidos em expressoes
}
