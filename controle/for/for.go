// AULA_26 - FOR
package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	fmt.Println("---while like")
	for i <= 10 { // nao tem while em go
		fmt.Print(i, ", ")
		i++
	}

	fmt.Println("\n---for")
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d, ", i)
	}

	fmt.Println("\n---Misturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("impar ")
		}
	}

	fmt.Println("\n---laco infinito")
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}

	// veremos o foreach no capitulo de array

}
