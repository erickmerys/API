package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/repostas"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
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

	if erro = usuario.Preparar("cadastro"); erro != nil {
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
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuario, erro := repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario altera os dados de usuário dentro do banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuario"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		repostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edição"); erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario apaga um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusNoContent, nil)
}
