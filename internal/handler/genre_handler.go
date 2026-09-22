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

func (g *GenreHandler) GetGenres(ctx *gin.Context) {
	genres, err := g.gs.GetGenres(ctx.Request.Context())
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	genreResponse := make([]model.GenreResponse, 0, len(genres))

	for _, genre := range genres {
		genreResponse = append(genreResponse, model.GenreResponse{
			ID:        genre.ID,
			Name:      genre.Name,
			CreatedAt: genre.CreatedAt,
			UpdatedAt: genre.UpdatedAt,
		})
	}

	response.ResponseSuccess(ctx, http.StatusOK, "Genres retrieved successfully", genreResponse)
}
