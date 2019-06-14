package main

import (
	"fmt"
	"time"

	"github.com/rodrigoaraujoti/go_area/htmlutil"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := htmlutil.Titulo(url1)
	c2 := htmlutil.Titulo(url2)
	c3 := htmlutil.Titulo(url3)

	// estrutura de concorrencia especifica para concorrencia (switch para concorrencia)
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Second):
		return "Todos perderam!"
		// default:
		// 	return "Sem retorno ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		// "https://www.cod3r.com.br",
		"https://www.google.com",
		"https://g1.globo.com",
		"https://www.youtube.com",
	)
	fmt.Println(campeao)
}
