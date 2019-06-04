// AULA_016 - OPERADORES ARITMETICOS
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao => ", a-b)
	fmt.Println("Multiplicacao => ", a*b)
	fmt.Println("Divisao => ", a/b)
	fmt.Println("Módulo =>", a%b)

	//bitwise - bit a bit
	fmt.Println("E =>", a&b)    // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)   // 11 | 10 = 11
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operacoes usando math
	// fmt.Println("Maior => ", math.Max(a, b)) //>>cannot use a (type int) as type float64 in argument to math.Max
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Exponenciacao => ", math.Pow(c, d))
	// math...

}
