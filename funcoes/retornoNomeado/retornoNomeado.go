// AULA_042
package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo - os valores ja foram atribuidos nas variaveis de retorno
}

func trocarSemRetornoNomeado(p1, p2 int) (int, int) {
	return p2, p1
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
