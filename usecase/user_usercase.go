package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UsuarioUseCase struct {
	respository repository.UsuarioRepository
}

func NewUsuarioUseCase(repositoryUsu repository.UsuarioRepository) UsuarioUseCase {
	return UsuarioUseCase{respository: repositoryUsu}
}

func (p *UsuarioUseCase) GetUsuarios() ([]model.Usuarios, error) {
	return p.respository.GetUsuarios()
}

func (p *UsuarioUseCase) CreateUsuario(NewUsuario model.Usuarios) error {
	return p.respository.CreateUsuario(NewUsuario)
}
