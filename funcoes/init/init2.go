// AULA_052
// - este arquivo eh uma copia do init.go
// no terminal, ir neste diretorio e executar: > go run *.go

/* - se houver dois mains no pacote:
$ go run *.go
# command-line-arguments
.\init2.go:10:6: main redeclared in this block
        previous declaration at .\init.go:10:6
*/
package main

import "fmt"

func init() {
	fmt.Println("Inicializando2... ")
}

/* func main() {
	fmt.Println("Main... ")
} */
