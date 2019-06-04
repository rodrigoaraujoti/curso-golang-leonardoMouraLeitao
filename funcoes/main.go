package main

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}

// go run main.go -> erros de compilacao:
// .\main.go:4:15: undefined: somar
// .\main.go:5:2: undefined: imprimir

// go run *.go -> roda tudo o que esta no pacote
// , a partir da main, e os metodos em outros arquivos sao reconhecidos
