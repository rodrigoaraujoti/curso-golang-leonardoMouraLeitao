// DESAFIO PROPOSTO NA AULA 028
// REESCREVER O EXEMPLO DO IFELSEIF USANDO O SWITCH TRUE
package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n > 10 || n < 0:
		return "Valor inválido!"
	case n >= 9:
		return "A"
	case n >= 8:
		return "B"
	case n >= 5:
		return "C"
	case n >= 3:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(11.1))
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.3))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(4.5))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(-11.1))
}
