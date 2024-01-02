package main

import (
	"log"
	"mysql/server"
	
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/criar", server.CriarUsuario)
	router.GET("/usuarios", server.BuscarUsuarios)
	router.GET("/usuarios/:id", server.BuscarUsuario)
	router.PUT("/usuarios/:id", server.AtualizarUsuario)
	router.DELETE("/usuarios/:id", server.DeletarUsuario)

	log.Fatal(router.Run(":8000"))
}