// AULA_23
package main

import "fmt"

func imprimirResultado(nota float64) {
	// if (nota >= 7) { --> o formatador remove os parenteses que englobam toda a expressao
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
