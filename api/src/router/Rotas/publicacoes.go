package Rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rotas{
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoes,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/publicacoes",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.BuscarPublicacoesPorUsuario,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/curtir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CurtirPublicacao,
		RequestAutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/descurtir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.DescurtirPublicacao,
		RequestAutenticacao: true,
	},
}
