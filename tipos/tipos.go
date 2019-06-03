package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// numero inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro 32000 é", reflect.TypeOf(32000))

	// inteiros sem sinal (só positivos)... uint8 uint16 uint32 uint64 (em java: byte, short, int, long)
	// u - unsigned int
	var b byte = 255 //byte eh um ALIAS para o tipo uint8
	fmt.Println("O byte é", reflect.TypeOf(b))
	fmt.Println("O tipo de b é", reflect.TypeOf(b))

	// inteiros com sinal... int8 int16 int32 int64 (em java: byte, short, int, long)
	// positivos e negativos
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // rune representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rune eh", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	var x2 = float32(49.99)
	fmt.Println("O tipo do x eh", reflect.TypeOf(x))
	fmt.Println("O tipo do x2 eh", reflect.TypeOf(x2))
	fmt.Println("O tipo do literal 49.99 eh", reflect.TypeOf(49.99)) //float64

	// boolean
	// bo := 1 #invalido
	bo := true
	fmt.Println("O tipo de bo eh", reflect.TypeOf(bo))
	fmt.Println("O valor de bo eh", bo)
	fmt.Println("A negacao do valor de bo eh", !bo)

	// String
	s1 := "Ola meu nome eh Leo"
	fmt.Println(s1 + "!") // concatenando strings com +
	fmt.Println("O tamanho do string eh", len(s1))

	// String com multiplas linhas
	s2 := `Ola
	meu
	nome
	eh
	leo
	`
	fmt.Println(s2)

	// char??? nao eh um tipo, deve se usar o int32
	//var a char = 'a' #invalido
	char := 'a'
	fmt.Println("O tipo DA VARIAVEL char eh", reflect.TypeOf(char))
	fmt.Println("O valor da variavel char eh", char)

}
