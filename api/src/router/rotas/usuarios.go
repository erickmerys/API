package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotaUsuarios = []Rotas{
	//Esse metodo cadastra um usuario
	{
		URI:                 "/usuarios",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarUsuario,
		RequestAutenticacao: false,
	},
	{
		URI:                 "/usuarios",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuarios,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuario,
		RequestAutenticacao: false,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUsuario,
		RequestAutenticacao: false,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequestAutenticacao: false,
	},
}
