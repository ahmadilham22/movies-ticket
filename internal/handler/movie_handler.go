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

func (h *MovieHandler) GetMovies(ctx *gin.Context) {
	movies, err := h.ms.GetMovies(ctx.Request.Context())
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	movieResponses := make([]model.MovieResponse, 0, len(movies))

	for _, item := range movies {
		genreResponses := make([]model.GenreSummaryResponse, 0, len(item.Genres))

		for _, genre := range item.Genres {
			genreResponses = append(
				genreResponses,
				model.GenreSummaryResponse{
					ID:   genre.ID,
					Name: genre.Name,
				},
			)
		}

		movieResponses = append(movieResponses, model.MovieResponse{
			ID:              item.Movie.ID,
			Title:           item.Movie.Title,
			Synopsis:        item.Movie.Synopsis,
			DurationMinutes: item.Movie.DurationMinutes,
			ReleaseDate:     item.Movie.ReleaseDate.Format(time.DateOnly),
			PosterURL:       item.Movie.PosterURL,
			TrailerURL:      item.Movie.TrailerURL,
			AgeRating:       item.Movie.AgeRating,
			Language:        item.Movie.Language,
			Country:         item.Movie.Country,
			Genres:          genreResponses,
			CreatedAt:       item.Movie.CreatedAt,
			UpdatedAt:       item.Movie.UpdatedAt,
		})
	}
	response.ResponseSuccess(
		ctx,
		http.StatusOK,
		"Movies retrieved successfully",
		movieResponses,
	)
}

func (h *MovieHandler) GetMovieByID(ctx *gin.Context) {
	var req model.GetMovieByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid movie ID",
		})
		return
	}

	movie, err := h.ms.GetMoviesByID(ctx.Request.Context(), req.ID)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}

	genreResponses := make([]model.GenreSummaryResponse, 0, len(movie.Genres))

	for _, genre := range movie.Genres {
		genreResponses = append(
			genreResponses,
			model.GenreSummaryResponse{
				ID:   genre.ID,
				Name: genre.Name,
			},
		)
	}

	movieResponse := model.MovieResponse{
		ID:              movie.Movie.ID,
		Title:           movie.Movie.Title,
		Synopsis:        movie.Movie.Synopsis,
		DurationMinutes: movie.Movie.DurationMinutes,
		ReleaseDate:     movie.Movie.ReleaseDate.Format(time.DateOnly),
		PosterURL:       movie.Movie.PosterURL,
		TrailerURL:      movie.Movie.TrailerURL,
		AgeRating:       movie.Movie.AgeRating,
		Language:        movie.Movie.Language,
		Country:         movie.Movie.Country,
		Genres:          genreResponses,
		CreatedAt:       movie.Movie.CreatedAt,
		UpdatedAt:       movie.Movie.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Movie retrieved successfully",
		"data":    movieResponse,
	})

}
