// AULA_89 - CRIANDO UM SERVIDOR ESTATICO
package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) // dir a ser criado dentro do dir deste arquivo static.go
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// executar pelo terminal (no dir do arquivo static.go):
// go run static.go
