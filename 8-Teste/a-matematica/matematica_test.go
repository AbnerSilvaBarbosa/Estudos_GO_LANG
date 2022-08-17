package amatematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

func TestMedia(t *testing.T) {
	valorEsperado := 7.275
	valor := media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

const valor = 10

func TestValor(t *testing.T) {
	if valor != 10 {
		t.Errorf("Diferente de 10")
	}
}
