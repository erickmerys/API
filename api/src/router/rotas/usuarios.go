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
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.SeguirUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.PararDeSeguirUsuario,
		RequestAutenticacao: true,
	},
}
