// AULA 090 - GERANDO CONTEUDO DINAMICO
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") //01 mes, 02 dia, 03 hora, 04 minuto, 05 segundo, 06 ano
	fmt.Fprintf(w, "<H1>Hora certa: %s</H1>", s)

}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
