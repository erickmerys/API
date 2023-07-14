package router

import (
	"api/src/router/Rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return Rotas.Configurar(r)
}