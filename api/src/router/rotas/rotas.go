package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas para API
type Rotas struct {
	URI                 string
	Metodo              string
	Funcao              func(http.ResponseWriter, *http.Request)
	RequestAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotaUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
