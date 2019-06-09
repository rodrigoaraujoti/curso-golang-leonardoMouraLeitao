// AULA 50 - DEFER
// - a linha só sera executada no fim do bloco de codigo atual
// - muito utilizado para fechar recursos (como conexao com banco de dados)
package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	defer fmt.Println("Fim2!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
