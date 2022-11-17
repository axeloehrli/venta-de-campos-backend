package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/axeloehrli/venta-de-campos-backend/token"
	"github.com/axeloehrli/venta-de-campos-backend/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	db         *sql.DB
	tokenMaker token.Maker
	router     *gin.Engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewServer(config util.Config, db *sql.DB) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}
	server := &Server{
		config:     config,
		db:         db,
		tokenMaker: tokenMaker,
	}

	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/usuarios", server.createUsuario)
	router.POST("/usuarios/login", server.loginUsuario)
	router.GET("/campos", server.listCampos)
	router.GET("/filtered-campos", server.listFilteredCampos)
	router.GET("/filtered-campos-count", server.getFilteredCamposCount)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/verify", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "success")
	})
	authRoutes.GET("/usuarios/:nombre_usuario", server.getUsuario)
	authRoutes.GET("/usuarios", server.listUsuarios)
	authRoutes.DELETE("/usuarios/:id", server.deleteUsuario)

	authRoutes.POST("/campos", server.createCampo)
	authRoutes.GET("/campos/:id", server.getCampo)

	authRoutes.DELETE("/campos/:id", server.deleteCampo)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run("localhost:8000")
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
