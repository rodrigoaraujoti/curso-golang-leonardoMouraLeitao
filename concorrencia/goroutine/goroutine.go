// AULA_68
package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// execucao sem concorrencia
	// fale("Maria", "Pq vc nao fala comigo?", 3)
	// fale("Joao", "So posso falar depois de voce", 1)

	// com concorrencia, entretanto a main termina antes das goroutines, que acabam sendo canceladas
	// go fale("Maria", "Ei...", 500)
	// go fale("Joao", "Opa...", 500)
	// time.Sleep(time.Second * 10)

	go fale("Maria", "Entendi!!!", 10)
	fale("Joao", "Parabéns!", 5)

	fmt.Println("Fim!")

}
