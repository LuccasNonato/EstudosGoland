package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type UsuarioRepository struct {
	connectioDB *sql.DB
}

func NewUsuarioRepository(connectioDB *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{connectioDB: connectioDB}
}

func (p *UsuarioRepository) GetUsuarios() ([]model.Usuarios, error) {
	query := "SELECT id, username, senha FROM usuarios"

	rows, err := p.connectioDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuariosList []model.Usuarios
	for rows.Next() {
		var usuario model.Usuarios
		if err := rows.Scan(&usuario.ID, &usuario.Usuario, &usuario.Senha); err != nil {
			return nil, err
		}
		usuariosList = append(usuariosList, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuariosList, nil
}

func (p *UsuarioRepository) CreateUsuario(NewUsuario model.Usuarios) error {

	if NewUsuario.Usuario == "" || NewUsuario.Senha == "" {
		return fmt.Errorf("Informações não podem ser em branco")
	}

	query := "INSERT INTO usuarios (username, senha) VALUES ($1, $2) RETURNING id"
	_, err := p.connectioDB.Exec(query, NewUsuario.Usuario, NewUsuario.Senha)
	return err
}
