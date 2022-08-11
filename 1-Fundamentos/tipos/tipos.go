package main

import (
	f "fmt"
	m "math"
	r "reflect"
)

func main() {
	//numeros inteiros
	f.Println(1, 2, 1000)
	f.Println(r.TypeOf(1))

	//sem sinal (só positivo).. uint8(byte) uint16 uint32 uint64

	i1 := m.MaxInt64

	f.Println("O valor maximo do int é", i1)
	f.Println("O tipo de i1 é", r.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	f.Println("O rune é ", r.TypeOf(i2))
	f.Println(i2)

	// números reais (float32, float64)

	//boolean
	bo := true
	f.Println("O tipo de bo é", r.TypeOf(bo))

	//String
	s1 := "Olá meu nome é Abner"
	f.Println(s1 + "!")
	f.Println("O tamanho da string é", len(s1))
}
