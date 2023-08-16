package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/repostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
)

// CriarUsuario cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
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

	usuarioaIDNoToken, erro := autenticacao.ExtrairUsuaripID(r)
	if erro != nil {
		repostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioaIDNoToken {
		repostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu!"))
		return
	}

	corpoRequest, erro := io.ReadAll(r.Body)
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

	usuarioaIDNoToken, erro := autenticacao.ExtrairUsuaripID(r)
	if erro != nil {
		repostas.Erro(w, http.StatusUnauthorized, erro)
	}

	if usuarioID != usuarioaIDNoToken {
		repostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o seu!"))
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

// SeguirUsuario permiti que um usuário siga outro
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuaripID(r)
	if erro != nil {
		repostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		repostas.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro = repositorio.Seguir(usuarioID, seguidorID); erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario permite que um usuário deixe de seguir outro
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuaripID(r)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == usuarioID {
		repostas.Erro(w, http.StatusForbidden, errors.New("Não é possível parar de seguir você mesmo"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	if erro = repositorio.ParaDeSeguir(usuarioID, seguidorID); erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarSeguidores traz todos os seguidores de um usuário
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	UsuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	seguidores, erro := repositorio.BuscarSeguidores(UsuarioID)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)
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
	usuarios, erro := repositorio.BuscarSeguindo(usuarioID)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha permite atualizar uma senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioaIDNoToken, erro := autenticacao.ExtrairUsuaripID(r)
	if erro != nil {
		repostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametro := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametro["usuarioId"], 10, 64)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioaIDNoToken != usuarioID {
		repostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de um usuário que não seja o seu"))
	}

	corpoRequest, erro := io.ReadAll(r.Body)

	var senha modelos.Senha
	if erro = json.Unmarshal(corpoRequest, &senha); erro != nil {
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
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(senhaSalvaNoBanco, senha.Atual); erro != nil {
		repostas.Erro(w, http.StatusUnauthorized, errors.New("A senha atual não condiz com a que está salva no banco"))
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		repostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		repostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repostas.JSON(w, http.StatusOK, nil)
}
