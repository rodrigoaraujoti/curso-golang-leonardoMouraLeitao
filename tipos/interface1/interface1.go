// AULA_058 - USANDO INTERFACES #01
// - Em GO nao se declara que um "tipo" implementa uma interface,
// 		se o "tipo" tem os metodos da interface, automaticamente ele implementa a interface
// - Forma como o GO lang possui polimorfismo
package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calca Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)
	imprimir(produto{"Calca Jeans", 79.90})

	p2 := produto{"Lapis", 1.50}
	imprimir(p2)
}
