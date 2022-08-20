package main

import (
	"log"
	"net/http"
)

// Não e possivel iniciar o servidor pela extensão runner/ cmd: cd a-static
// cmd: go run static.go

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	http.ListenAndServe(":3000", nil)
}
