// AULA_021 - NAO EXISTE OPERADOR TERNARIO
package main

import "fmt"

// nao tem operador ternario
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
	// return nota > 6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
