package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotaLogin = Rotas{
	URI:                 "/login",
	Metodo:              http.MethodPost,
	Funcao:              controllers.Login,
	RequestAutenticacao: false,
}
