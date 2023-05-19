package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rotas struct {
URI					string
Metodo				string
Funcao				func(http.ResponseWriter, *http.Request)
RequestAutenticacao	bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	
	return r
}