package repository

import (
	"context"
	"errors"
	"online-ticketing/internal/model"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

var ErrGenreAlreadyExists = errors.New("genre already exists")

type GenreRepository struct {
	db *sqlx.DB
}

func NewGenreRepository(db *sqlx.DB) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

func (g *GenreRepository) CreateGenre(ctx context.Context, name string) (model.Genre, error) {
	var genre model.Genre
	query := `INSERT INTO genres (name) VALUES ($1) RETURNING id, name, created_at, updated_at`
	err := g.db.GetContext(ctx, &genre, query, name)
	if err != nil {
		if pgError, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgError.Code == "23505" && pgError.ConstraintName == "uidx_genres_name_ci" {
				return genre, ErrGenreAlreadyExists
			}
		}
		return genre, err
	}

	return genre, nil

}

func (g *GenreRepository) GetGenres(ctx context.Context) ([]model.Genre, error) {
	genres := []model.Genre{}
	query := "SELECT id, name, created_at, updated_at FROM genres ORDER BY LOWER(BTRIM(name)) ASC"
	err := g.db.SelectContext(ctx, &genres, query)
	if err != nil {
		return nil, err
	}

	return genres, nil
}
