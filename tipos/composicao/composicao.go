// AULA_060 - COMPOSICAO DE INTERFACES
package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais metodos
}

type bmw7 struct {
}

// func (e esportivoLuxuoso) ligarTurbo() {-> invalid receiver type esportivoLuxuoso (esportivoLuxuoso is an interface type)go
func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()

	// b2 := esportivo{} -> invalid type for composite literal: esportivo
}
