// AULA_063 - PACOTES E VISIBILIDADE
package main

import "math"

// inicializando com letra maiuscula eh PUBLICO (visivel dentro e fora do pacote)!
// inicializando com letra minuscula eh PRIVADO (visivel apenas dentro do pacote)!
//		quando compilado, todos os arquivos se transformam

// por exemplo...
// Ponto - gerara algo publico
// ponto ou _Ponto - gerara algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return // retorno limpo - devido ao retorno nomeado na assinatura do metodo
}

// Distancia eh o Calculo da distancia linear entre os dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
