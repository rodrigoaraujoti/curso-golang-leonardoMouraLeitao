// AULA043
package main

import "fmt"

// soma := func(a, b int) int { ->syntax error: non-declaration statement outside function bodygo
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
