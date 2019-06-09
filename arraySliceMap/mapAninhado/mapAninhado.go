// AULA_039 - MAPS ANINHADOS
package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"Jose Joao": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	println("\nexercicio: printar maps aninhados")
	for letra, funcs := range funcsPorLetra {
		fmt.Println("- ", letra)
		for funcionario, salario := range funcs {
			fmt.Println(funcionario, salario)
		}
	}

	fmt.Println("\nprint de objeto map aninhado\n", funcsPorLetra)

}
