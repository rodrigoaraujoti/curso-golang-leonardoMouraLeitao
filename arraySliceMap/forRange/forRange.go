// AULA_032 percorrendo arrays com for (range)
package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // o compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	fmt.Println("Exemplo de for range, com iterador ignerado, ao usar underscore")
	for _, num := range numeros {
		fmt.Println(num)
	}
}
