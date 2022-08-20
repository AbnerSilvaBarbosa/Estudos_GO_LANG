package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuario/", UsuarioHandler)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
