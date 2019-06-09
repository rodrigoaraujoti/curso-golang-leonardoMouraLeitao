// AULA_032 percorrendo arrays com for (range)
package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // o compilador conta!

	fmt.Println("for index, item := range arraySliceMap")
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	fmt.Println("Exemplo de for range, com iterador ignerado, ao usar underscore")
	fmt.Println("- O primeiro campo deste for eh sempre o indice, se ele nao for ser usado, informar o underscore")
	for _, num := range numeros {
		fmt.Println(num)
	}

	fmt.Println("- O segundo campo deste for eh sempre o item do range, se ele nao for ser usado, nao eh necessario informa-lo")
	for i := range numeros {
		fmt.Println(i)
	}
}
