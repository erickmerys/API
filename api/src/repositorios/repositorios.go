package repositorios

import (
	"api/src/modelos"
	"api/src/repostas"
	"database/sql"
	"fmt"
	"net/http"
)

// Usuario representa um repositorio de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um novo repositorio de usuário
func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios(nome, nick, email, senha) VALUES(?,?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio Usuarios) Buscar(usuarioOuNick string) ([]modelos.Usuario, error) {
	usuarioOuNick = fmt.Sprintf("%%%s%%", usuarioOuNick)

	linhas, erro := repositorio.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
	usuarioOuNick, usuarioOuNick); 
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		);erro != nil {
			return nil, erro
		}
		
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
