package Rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rotas{
	URI:                 "/login",
	Metodo:              http.MethodPost,
	Funcao:              controllers.Login,
	RequestAutenticacao: false,
}
