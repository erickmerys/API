package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")
	r:= router.Gerar()

// Informar o porta e a rota
	log.Fatal(http.ListenAndServe(":5000", r))
}