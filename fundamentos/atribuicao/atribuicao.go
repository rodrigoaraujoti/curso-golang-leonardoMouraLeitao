// AULA_017 - ATRIBUICAO DE TIPOS
package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2 // modulo
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x // atribuicao multipla: trocando valores de variáveis, sem precisar de uma terceira var temporaria
	fmt.Println(x, y)

	w, z, a := x, y, "a" // atribuicao multipla combinada com inferencia de tipo
	fmt.Println(w, z, a)

}
