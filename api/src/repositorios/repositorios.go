package repositorios

import (
	"api/src/modelos"
	"database/sql"
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
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios(nome, nick, email, senha) VALUES(?,?,?,?)",)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil

}	