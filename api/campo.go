package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/axeloehrli/venta-de-campos-backend/db"
	"github.com/axeloehrli/venta-de-campos-backend/util"
	"github.com/gin-gonic/gin"
)

type createCampoRequest struct {
	IDUsuario         int64  `json:"id_usuario" binding:"required"`
	Titulo            string `json:"titulo" binding:"required"`
	Descripcion       string `json:"descripcion" binding:"required"`
	Tipo              string `json:"tipo" binding:"required"`
	Hectareas         int64  `json:"hectareas" binding:"required,min=1"`
	PrecioPorHectarea int64  `json:"precio_por_hectarea" binding:"required,min=1"`
	Ciudad            string `json:"ciudad" binding:"required"`
	Provincia         string `json:"provincia" binding:"required"`
}

func (server *Server) createCampo(ctx *gin.Context) {
	var req createCampoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateCampoParams{
		IDUsuario:         req.IDUsuario,
		Titulo:            req.Titulo,
		Descripcion:       req.Descripcion,
		Tipo:              req.Tipo,
		Hectareas:         req.Hectareas,
		PrecioPorHectarea: req.PrecioPorHectarea,
		Ciudad:            req.Ciudad,
		Provincia:         req.Provincia,
	}

	c, err := db.CreateCampo(context.Background(), arg, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, c)
}

type getCampoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCampo(ctx *gin.Context) {
	var req getCampoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	c, err := db.GetCampo(context.Background(), req.ID, server.db)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, c)
}

type deleteCampoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCampo(ctx *gin.Context) {
	var req deleteCampoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := db.DeleteCampo(context.Background(), req.ID, server.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
}

type listCamposRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCampos(ctx *gin.Context) {
	var req listCamposRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	arg := db.ListCamposParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	c, err := db.ListCampos(context.Background(), arg, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, c)
}

type getFilteredCamposCountRequest struct {
	Provincia            *string `form:"provincia"`
	Tipo                 *string `form:"tipo"`
	PrecioPorHectareaMin *int64  `form:"precio_por_hectarea_min"`
	PrecioPorHectareaMax *int64  `form:"precio_por_hectarea_max"`
	HectareasMin         *int64  `form:"hectareas_min"`
	HectareasMax         *int64  `form:"hectareas_max"`
}

func (server *Server) getFilteredCamposCount(ctx *gin.Context) {
	var req getFilteredCamposCountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	filtersMap := util.FiltersStructToMap(req)

	conditions := util.BuildConditionsString(filtersMap)
	dynamicQuery := fmt.Sprintf("SELECT COUNT(*) FROM CAMPOS %v", conditions)
	c, err := db.GetFilteredCamposCount(context.Background(), dynamicQuery, server.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, c)
}

type listFilteredCamposRequest struct {
	PageID               int32   `form:"page_id" binding:"required,min=1"`
	PageSize             int32   `form:"page_size" binding:"required,min=5,max=10"`
	Provincia            *string `form:"provincia"`
	Tipo                 *string `form:"tipo"`
	PrecioPorHectareaMin *int64  `form:"precio_por_hectarea_min"`
	PrecioPorHectareaMax *int64  `form:"precio_por_hectarea_max"`
	HectareasMin         *int64  `form:"hectareas_min"`
	HectareasMax         *int64  `form:"hectareas_max"`
}

func (server *Server) listFilteredCampos(ctx *gin.Context) {
	var req listFilteredCamposRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	filtersMap := util.FiltersStructToMap(req)

	conditions := util.BuildConditionsString(filtersMap)
	dynamicQuery := fmt.Sprintf("SELECT * FROM CAMPOS %v LIMIT %v OFFSET %v", conditions, req.PageSize, (req.PageID-1)*req.PageSize)
	c, err := db.ListFilteredCampos(context.Background(), dynamicQuery, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, c)
}
