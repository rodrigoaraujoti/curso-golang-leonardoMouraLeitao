// AULA_014 - CONVERSOES
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 //default float64
	y := 2
	// fmt.Println(x / y) // invalid operation: x / y (mismatched types float64 and int)
	fmt.Println(x / float64(y)) // converteu o y para evitar o erro da linha acima
	fmt.Println(int(x) / y)     // outra forma de conversao para evitar o erro, mas perde a precisao do float

	nota := 6.9
	notaFinal := int(nota) // converter de float p int, apenas remove os decimais, sem arredondamento
	fmt.Println(notaFinal)

	// cuidado: conversao para string com numeros converte para codigo da tabela ascii
	fmt.Println("Teste " + string(97))
	// inteiro para string
	fmt.Println("Teste " + strconv.Itoa(123))
	// fmt.Println("Teste " + fmt.Stringer(123)) fmt.sstring?->pesquisar, foi apenas citado

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // apenas "true" e "1" imprimirao "verdadeiro"
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}

}
