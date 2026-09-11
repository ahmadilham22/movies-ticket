package service

import (
	"context"
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
	"strings"
	"unicode/utf8"
)

var ErrGenreNameRequired = errors.New("genre name is required")
var ErrGenreNameTooLong = errors.New("genre name is too long")
var ErrGenreAlreadyExists = errors.New("genre name already exists")

type genreRepository interface {
	CreateGenre(ctx context.Context, name string) (model.Genre, error)
}

type GenreService struct {
	gr genreRepository
}

func NewGenreService(gr genreRepository) *GenreService {
	return &GenreService{
		gr: gr,
	}
}

func (g *GenreService) CreateGenre(ctx context.Context, name string) (model.Genre, error) {
	normalizedName := strings.TrimSpace(name)
	if normalizedName == "" {
		return model.Genre{}, ErrGenreNameRequired
	}

	if utf8.RuneCountInString(normalizedName) > 64 {
		return model.Genre{}, ErrGenreNameTooLong
	}

	result, err := g.gr.CreateGenre(ctx, normalizedName)

	if err != nil {
		if errors.Is(err, repository.ErrGenreAlreadyExists) {
			return model.Genre{}, ErrGenreAlreadyExists
		}
		return model.Genre{}, err
	}

	return result, nil

}
