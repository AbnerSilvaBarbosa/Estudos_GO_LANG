package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func horaCerta(res http.ResponseWriter, req *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	b, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()

	fmt.Fprintln(res, string(b))
	fmt.Println(string(b))

	fmt.Fprintf(res, "<h1>Hora certa: %s<h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	http.ListenAndServe(":3000", nil)
}
