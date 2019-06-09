// AULA_31 ARRAY
package main

import "fmt"

func main() {
	// homogenea (mesmo tipo) e estatica (fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10 -> invalid array index 3 (out of bounds for 3-element array)

	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ { // len() tamanho de array
		total += notas[i]
	}

	media := total / float64(len(notas)) // em divisao eh necessario que os dois valores sejam de um mesmo tipo
	fmt.Printf("Media %.2f\n", media)

}
