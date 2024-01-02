package server

import (
	"fmt"
	"mysql/banco"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(c *gin.Context) {
	var usuario usuario

	if e := c.ShouldBindJSON(&usuario); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao converter o usuário para struct!"})
		return
	}

	db, e := banco.Conectar()
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados!"})
		return
	}
	defer db.Close()

	result, e := db.Exec("INSERT INTO usuarios (nome, email) VALUES (?, ?)", usuario.Nome, usuario.Email)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao executar o statement!"})
		return
	}

	idInserido, e := result.LastInsertId()
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter o ID inserido!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)})
}

func BuscarUsuarios(c *gin.Context) {
	db, e := banco.Conectar()
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados!"})
		return
	}
	defer db.Close()

	rows, e := db.Query("SELECT * FROM usuarios")
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar os usuários!"})
		return
	}
	defer rows.Close()

	var usuarios []usuario

	for rows.Next() {
		var u usuario
		if e := rows.Scan(&u.ID, &u.Nome, &u.Email); e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao escanear o usuário!"})
			return
		}
		usuarios = append(usuarios, u)
	}

	c.JSON(http.StatusOK, usuarios)
}

func BuscarUsuario(c *gin.Context) {
	db, e := banco.Conectar()
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados!"})
		return
	}
	defer db.Close()

	var usuario usuario
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido!"})
		return
	}

	e = db.QueryRow("SELECT * FROM usuarios WHERE id = ?", id).Scan(&usuario.ID, &usuario.Nome, &usuario.Email)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar o usuário!"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func AtualizarUsuario(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido!"})
		return
	}

	var usuario usuario
	if e := c.ShouldBindJSON(&usuario); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao converter o usuário para struct!"})
		return
	}

	db, e := banco.Conectar()
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados!"})
		return
	}
	defer db.Close()

	_, e = db.Exec("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?", usuario.Nome, usuario.Email, id)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o usuário!"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Usuário atualizado com sucesso!"})
}

func DeletarUsuario(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido!"})
		return
	}

	db, e := banco.Conectar()
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar com o banco de dados!"})
		return
	}
	defer db.Close()

	_, e = db.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o usuário!"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Usuário deletado com sucesso!"})
}
