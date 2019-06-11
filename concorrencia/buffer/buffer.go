// AULA_72 - CHANNEL COM BUFFER
package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou4!") //~ apos a quarta escrita ainda ocasionalmente o "executou" eh printado
	ch <- 5
	fmt.Println("Executou5!") //~ já na quinta linha isto nao ocorreu
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
