package main

import (
	controller "go-api/Controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	UsuarioRepository := repository.NewUsuarioRepository(dbConnection)
	//camada dos usecases
	UsuarioUsecase := usecase.NewUsuarioUseCase(*UsuarioRepository)

	//camada dos controllers
	UsuariosController := controller.NewUsuarioController(UsuarioUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/usuarios", UsuariosController.GetUsuarios)
	server.POST("/create-usuario", UsuariosController.CreateUsuario)
	server.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
