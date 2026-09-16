package handler

import (
	"net/http"
	"online-ticketing/internal/model"
	"online-ticketing/internal/response"
	"online-ticketing/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	ms *service.MovieService
}

func NewMovieHandler(ms *service.MovieService) *MovieHandler {
	return &MovieHandler{ms: ms}
}

func (h *MovieHandler) CreateMovie(ctx *gin.Context) {
	movieRequest := model.CreateMovieRequest{}

	if err := ctx.ShouldBindJSON(&movieRequest); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	movie, genres, err := h.ms.CreateMovie(ctx.Request.Context(), movieRequest)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	genreResponses := make([]model.GenreSummaryResponse, 0, len(genres))
	for _, genre := range genres {
		genreResponses = append(genreResponses, model.GenreSummaryResponse{
			ID:   genre.ID,
			Name: genre.Name,
		})
	}

	result := model.MovieResponse{
		ID:              movie.ID,
		Title:           movie.Title,
		Synopsis:        movie.Synopsis,
		DurationMinutes: movie.DurationMinutes,
		ReleaseDate:     movie.ReleaseDate.Format(time.DateOnly),
		PosterURL:       movie.PosterURL,
		TrailerURL:      movie.TrailerURL,
		AgeRating:       movie.AgeRating,
		Language:        movie.Language,
		Country:         movie.Country,
		Genres:          genreResponses,
		CreatedAt:       movie.CreatedAt,
		UpdatedAt:       movie.UpdatedAt,
	}

	response.ResponseSuccess(
		ctx,
		http.StatusCreated,
		"Movie created successfully",
		result,
	)

}
