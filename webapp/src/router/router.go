package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)


func Gerar() *mux.Router {
	r := mux.NewRouter()
	rotas.Configurar(r)
	return r
}