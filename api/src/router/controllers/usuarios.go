package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/repostas"
	"encoding/json"
	"io/ioutil"
	"strings"

	"net/http"
)

// CriarUsuario cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		repostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := usuario.Preparar(); erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca mais de um usuário do banco dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarioOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarios, erro := repositorio.Buscar(usuarioOuNick)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusOK, usuarios)
}

// BusacarUsuario busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário em especifico!"))
}

// AtualizarUsuario altera os dados de usuário dentro do banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualização de usuário completa!"))
}

// DeletarUsuario apaga um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
