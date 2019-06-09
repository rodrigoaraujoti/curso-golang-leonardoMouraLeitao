// AULA_037 - MAPS #1
package main

import "fmt"

func main() {
	// var aprovadosErroExecucao map[int]string ->erroDeExecucao: panic: assignment to entry in nil map
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"
	aprovados[789] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[789]) // lendo um item do map usando a sua chave
	delete(aprovados, 789)      // removendo um item do map
	fmt.Println("lendo item deletado: ", aprovados[789])
}
