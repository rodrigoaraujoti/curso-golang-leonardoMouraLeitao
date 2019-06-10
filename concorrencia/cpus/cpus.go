// AULA_67 - NUMERO DE CPU'S
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
