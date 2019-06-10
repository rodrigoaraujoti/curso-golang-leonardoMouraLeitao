// AULA_058 - USANDO INTERFACES #01
// - Interface sao comumente utilizadas para se ler valores
package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual uint
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// ao usar o tipo da interface
	// e devido ao metodo receiver possuir atributo com ponteiro
	// eh necessario adicionar o &
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
