package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplate()

	fmt.Println("Escutando na porta 2000!")
	log.Fatal(http.ListenAndServe(":2000", router.Gerar()))
}
