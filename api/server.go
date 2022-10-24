package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	router *gin.Engine
}

func NewServer(db *sql.DB) *Server {
	server := &Server{db: db}

	router := gin.Default()
	router.POST("/usuarios", server.createUsuario)
	router.GET("/usuarios/:id", server.getUsuario)
	router.GET("/usuarios", server.listUsuarios)
	router.DELETE("/usuarios/:id", server.deleteUsuario)

	router.POST("/campos", server.createCampo)
	router.GET("/campos/:id", server.getCampo)
	router.GET("/campos", server.listCampos)
	router.DELETE("/campos/:id", server.deleteCampo)

	server.router = router

	return server
}

func (server *Server) Start() error {
	return server.router.Run("localhost:8000")
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
