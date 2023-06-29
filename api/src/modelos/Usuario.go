package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa o um usuário acessando uma rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome do usuário não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick do usuário não poder estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email não pode estar em branco")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar(etapa)
	return nil
}

func (usuario *Usuario) formatar(etapa string) {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}