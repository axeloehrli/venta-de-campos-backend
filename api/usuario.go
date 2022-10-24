package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/axeloehrli/venta-de-campos-backend/db"
	"github.com/gin-gonic/gin"
)

type createUsuarioRequest struct {
	NombreUsuario string `json:"nombre_usuario" binding:"required"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
}

func (server *Server) createUsuario(ctx *gin.Context) {
	var req createUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateUsuarioParams{
		NombreUsuario: req.NombreUsuario,
		Nombre:        req.Nombre,
		Apellido:      req.Apellido,
		Email:         req.Email,
	}

	u, err := db.CreateUsuario(context.Background(), arg, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, u)
}

type getUsuarioRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUsuario(ctx *gin.Context) {
	var req getUsuarioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	u, err := db.GetUsuario(context.Background(), req.ID, server.db)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, u)
}

type deleteUsuarioRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteUsuario(ctx *gin.Context) {
	var req deleteUsuarioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := db.DeleteUsuario(context.Background(), req.ID, server.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
}

type listUsuariosRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUsuarios(ctx *gin.Context) {
	var req listUsuariosRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	arg := db.ListUsuariosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	u, err := db.ListUsuarios(context.Background(), arg, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, u)
}
