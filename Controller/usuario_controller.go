package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	UsuarioUseCase usecase.UsuarioUseCase
}

func NewUsuarioController(usecase usecase.UsuarioUseCase) UsuarioController {
	return UsuarioController{
		UsuarioUseCase: usecase,
	}
}

func (p *UsuarioController) GetUsuarios(c *gin.Context) {
	usuario, err := p.UsuarioUseCase.GetUsuarios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (p *UsuarioController) CreateUsuario(c *gin.Context) {
	var newUsuario model.Usuarios
	if err := c.BindJSON(&newUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.UsuarioUseCase.CreateUsuario(newUsuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUsuario)
}
