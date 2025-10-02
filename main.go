package main

import (
	"cabother/aula/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Criar usuário
	router.POST("/user", handler.NewUser)
	router.POST("/user/addressCep/:number", handler.RandomCep)
	router.GET("/user/books", handler.GetAllUsersAndBooks)

	// criar um trabalho
	router.POST("/jobs", handler.NewJob)
	// Deletar Usuário
	//sair do trabalho
	router.DELETE("/jobs/:id", handler.RemoveJob)
	router.DELETE("/users", handler.RemoveUsersByLikeName)
	// Atualizar Usuário
	router.PUT("/users/:id", handler.UpdateUser)
	router.PUT("/jobs/:id", handler.UpdateJobByID)

	// Buscar Usuário
	router.GET("/userss/:id", handler.GetUsersByID)
	router.GET("/users", handler.GetAllUsers)
	// Buscar trabalho
	router.GET("/jobs/:id", handler.GetJobByID)
	router.GET("/jobs", handler.GetAllJobs)
	router.Run(":8080")
	router.DELETE("/users/:id", handler.RemoveUsers)
}
