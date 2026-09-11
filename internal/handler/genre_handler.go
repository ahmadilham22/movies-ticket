package handler

import (
	"net/http"
	"online-ticketing/internal/model"
	"online-ticketing/internal/response"
	"online-ticketing/internal/service"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	gs *service.GenreService
}

func NewGenreHandler(gs *service.GenreService) *GenreHandler {
	return &GenreHandler{
		gs: gs,
	}
}

func (g *GenreHandler) CreateGenre(ctx *gin.Context) {
	genreRequest := model.CreateGenreRequest{}

	if err := ctx.ShouldBindJSON(&genreRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	result, err := g.gs.CreateGenre(ctx.Request.Context(), genreRequest.Name)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	genreResponse := model.GenreResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	response.ResponseSuccess(ctx, http.StatusCreated, "Genre created successfully", genreResponse)
}
