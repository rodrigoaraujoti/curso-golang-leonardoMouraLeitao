// AULA_070 - USANDO GOROUTINE E CHANNEL
package main

import (
	"fmt"
	"time"
)

// Channel - eh a forma de comunicacao entre as goroutines
// channel eh um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o channel

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c //recebendo dados do canal (soh continua apos receber 2 dados)
	fmt.Println("B")

	fmt.Println(a, b)

	fmt.Println(<-c)
	// fmt.Println(<-c) ~fatal error: all goroutines are asleep - deadlock!
}
