// AULA_057 TIPO PERSONALIZADO
package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	// return float64(n) >= inicio && float64(n) <= fim
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10) {
		return "A"
	} else if n.entre(7, 8.99) {
		return "B"
	} else if n.entre(5, 6.99) {
		return "C"
	} else if n.entre(3, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func notaParaConceitoSwitch(n nota) string {
	switch {
	case n.entre(9.0, 10):
		return "A"
	case n.entre(7, 8.99):
		return "B"
	case n.entre(5, 6.99):
		return "C"
	case n.entre(3, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceitoSwitch(6.9))
	fmt.Println(notaParaConceito(2.1))
}
