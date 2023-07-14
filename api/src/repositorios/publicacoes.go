package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct{
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um novo repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	
}