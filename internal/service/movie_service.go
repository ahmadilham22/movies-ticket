package service

import (
	"context"
	"errors"
	"net/url"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	ErrInvalidMovieData        = errors.New("invalid movie data")
	ErrInvalidMovieDuration    = errors.New("invalid movie duration")
	ErrInvalidMovieReleaseDate = errors.New("invalid movie release date")
	ErrMovieGenresRequired     = errors.New("movie genres required")
	ErrDuplicateGenreIDs       = errors.New("duplicate genre IDs")
	ErrInvalidPosterURL        = errors.New("invalid poster URL")
	ErrInvalidTrailerURL       = errors.New("invalid trailer URL")
	ErrMovieGenreNotFound      = errors.New("movie genre not found")
)

func isValidHTTPURL(httpURL string) bool {
	parsed, err := url.ParseRequestURI(httpURL)
	if err != nil {
		return false
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}

	if parsed.Hostname() == "" {
		return false
	}

	return true
}

func exceedsRuneLimit(value string, max int) bool {
	result := utf8.RuneCountInString(value)
	return result > max
}

type movieRepository interface {
	CreateMovie(ctx context.Context, movie model.Movie, genreIDs []string) (model.Movie, []model.Genre, error)
}

type MovieService struct {
	repo movieRepository
}

func NewMovieService(repo movieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) CreateMovie(ctx context.Context, request model.CreateMovieRequest) (model.Movie, []model.Genre, error) {
	// Normalize input to trim spaces
	title := strings.TrimSpace(request.Title)
	synopsis := strings.TrimSpace(request.Synopsis)
	posterURL := strings.TrimSpace(request.PosterURL)
	releaseDateNormalized := strings.TrimSpace(request.ReleaseDate)
	ageRating := strings.TrimSpace(request.AgeRating)
	language := strings.TrimSpace(request.Language)
	country := strings.TrimSpace(request.Country)

	if title == "" || synopsis == "" || releaseDateNormalized == "" || posterURL == "" || ageRating == "" || language == "" || country == "" {
		return model.Movie{}, nil, ErrInvalidMovieData
	}

	if request.DurationMinutes <= 0 {
		return model.Movie{}, nil, ErrInvalidMovieDuration
	}

	releaseDate, err := time.Parse(time.DateOnly, releaseDateNormalized)
	if err != nil {
		return model.Movie{}, nil, ErrInvalidMovieReleaseDate
	}

	if len(request.GenreIDs) == 0 {
		return model.Movie{}, nil, ErrMovieGenresRequired
	}

	var trailerURL *string
	if request.TrailerURL != nil {
		trimmed := strings.TrimSpace(*request.TrailerURL)
		if trimmed != "" {
			trailerURL = &trimmed
		}
	}

	if exceedsRuneLimit(title, 255) {
		return model.Movie{}, nil, ErrInvalidMovieData
	}

	if exceedsRuneLimit(ageRating, 10) {
		return model.Movie{}, nil, ErrInvalidMovieData
	}

	if exceedsRuneLimit(language, 64) {
		return model.Movie{}, nil, ErrInvalidMovieData
	}

	if exceedsRuneLimit(country, 64) {
		return model.Movie{}, nil, ErrInvalidMovieData
	}

	seenGenreIDs := make(map[string]struct{}, len(request.GenreIDs))
	normalizedGenreIDs := make([]string, 0, len(request.GenreIDs))
	
	for _, genreID := range request.GenreIDs {
		normalizedID := strings.ToLower(strings.TrimSpace(genreID))
		_, exists := seenGenreIDs[normalizedID]
		if exists {
			return model.Movie{}, nil, ErrDuplicateGenreIDs
		}
		seenGenreIDs[normalizedID] = struct{}{}
		normalizedGenreIDs = append(normalizedGenreIDs, normalizedID)
	}

	if trailerURL != nil && !isValidHTTPURL(*trailerURL) {
		return model.Movie{}, nil, ErrInvalidTrailerURL
	}

	if !isValidHTTPURL(posterURL) {
		return model.Movie{}, nil, ErrInvalidPosterURL
	}

	movieRequest := model.Movie{
		Title:           title,
		Synopsis:        synopsis,
		ReleaseDate:     releaseDate,
		PosterURL:       posterURL,
		TrailerURL:      trailerURL,
		AgeRating:       ageRating,
		Language:        language,
		Country:         country,
		DurationMinutes: request.DurationMinutes,
	}

	movie, genres, err := s.repo.CreateMovie(ctx, movieRequest, normalizedGenreIDs)
	if err != nil {
		if errors.Is(err, repository.ErrGenreNotFound) {
			return model.Movie{}, nil, ErrMovieGenreNotFound
		}
		return model.Movie{}, nil, err
	}

	return movie, genres, nil
}
