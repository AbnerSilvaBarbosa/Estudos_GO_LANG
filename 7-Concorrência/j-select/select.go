package main

import (
	"fmt"
	"time"

	"github.com/AbnerSilvaBarbosa/area/area2"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1, c2, c3 := area2.Titulo(url1), area2.Titulo(url2), area2.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(5000 * time.Millisecond):
		return "Todos perderam!"
		//default:
		//	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido("https://www.youtube.com/", "https://github.com/", "https://stackoverflow.com/")
	fmt.Println(campeao)
}
