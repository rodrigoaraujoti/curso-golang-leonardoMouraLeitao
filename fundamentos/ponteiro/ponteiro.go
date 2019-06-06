// AULA 22 - Ponteiro
package main

func main() {
	i := 1
	// var j int = nil --> cannot use nil as type int in assignmentgo

	// GO nao tem aritmetica de ponteiros
	var p *int = nil

	p = &i // &var pega o ponteiro da variavel
	*p++   // desreferenciando (pegando o valor)
	i++    // incrementando variavel int (var que esta sendo referenciada pelo *p)
	// p++ --> nao desreferenciou o ponteiro, portanto nao permite usar operador unario

}

// ref: ponteiros, referenciar endereco de memoria vai de encontro com a programacao funcional
// ref: pop (place oriented program) programacao orientada a lugar, autor do clojure
