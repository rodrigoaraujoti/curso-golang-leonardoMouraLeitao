// AULA_051 - PASSANDO PONTEIRO PARA FUNCAO
// - tente evitar a passagem de referencias para outros metodos
// 		pois assim a funcao deixa de ser pura e assim pode editar coisas externas a ela
package main

import "fmt"

func inc1(n int) {
	n++
}

// revisao: um ponteiro eh representado por um *
func inc2(n *int) {
	// revisao: * eh usado para desreferenciar, ou seja
	// ter acesso ao valor ao qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n) // por valor
	fmt.Println(n)

	// revisao: & eh usado para obter o endereco da variavel
	inc2(&n) // por referencia
	fmt.Println(n)

}
