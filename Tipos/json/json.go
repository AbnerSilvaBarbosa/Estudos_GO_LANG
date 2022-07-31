package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	NOME  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{
		ID:    1,
		NOME:  "Notebook",
		Preco: 1890.0,
		Tags:  []string{"Promoção", "Eletrônico"},
	}

	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta"}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Print(p2.Preco)
}
