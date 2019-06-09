// AULA_O4* - RECURSIVIDADE
// (refatorada para receber apenas uint e assim nao precisar ter segundo retorno do tipo error)
package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
