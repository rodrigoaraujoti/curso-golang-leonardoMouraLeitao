// AULA-010 CONSTANTES E VARIAVEIS
package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	// forma reduzida de criar uma var - "declaracao..."
	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "Epa!"
	fmt.Println(g, h, i)

}
