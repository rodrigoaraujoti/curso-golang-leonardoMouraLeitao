// AULA038
package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, //ultimo item tambem tem que ter a virgula, ou a chave nesta mesma linha
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") // tentar deletar um item inexistente nao causa erro

	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios { // ?o range parece mudar a ordem dos itens a cada compilacao?
		fmt.Println(nome, salario)
	}
}
