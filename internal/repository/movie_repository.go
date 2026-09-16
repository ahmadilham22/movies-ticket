package repository

import (
	"context"
	"errors"
	"online-ticketing/internal/model"

	"github.com/jmoiron/sqlx"
)

var (
	ErrGenreNotFound            = errors.New("genre not found")
	ErrMovieGenreInsertMismatch = errors.New("movie genre insert mismatch")
)

type MovieRepository struct {
	db *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) CreateMovie(ctx context.Context, movie model.Movie, genreIDs []string) (model.Movie, []model.Genre, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Movie{}, nil, err
	}
	defer tx.Rollback()

	genres := []model.Genre{}

	query := "SELECT id, name, created_at, updated_at FROM genres WHERE id IN (?) ORDER BY LOWER(BTRIM(name)) FOR KEY SHARE"

	expandedQuery, args, err := sqlx.In(query, genreIDs)

	if err != nil {
		return model.Movie{}, nil, err
	}

	expandedQuery = tx.Rebind(expandedQuery)

	err = tx.SelectContext(ctx, &genres, expandedQuery, args...)

	if err != nil {
		return model.Movie{}, nil, err
	}

	if len(genres) != len(genreIDs) {
		return model.Movie{}, nil, ErrGenreNotFound
	}

	createdMovie := model.Movie{}

	movieQuery := "INSERT INTO movies (title, synopsis, duration_minutes, release_date, poster_url, trailer_url, age_rating, language, country) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, title, synopsis, duration_minutes, release_date, poster_url, trailer_url, age_rating, language, country, created_at, updated_at"

	err = tx.GetContext(ctx, &createdMovie, movieQuery, movie.Title, movie.Synopsis, movie.DurationMinutes, movie.ReleaseDate, movie.PosterURL, movie.TrailerURL, movie.AgeRating, movie.Language, movie.Country)

	if err != nil {
		return model.Movie{}, nil, err
	}

	genreQuery := "INSERT INTO movie_genres (movie_id, genre_id) SELECT ?, id FROM genres WHERE id IN (?)"

	expandedGenres, args, err := sqlx.In(genreQuery, createdMovie.ID, genreIDs)

	if err != nil {
		return model.Movie{}, nil, err
	}

	expandedGenres = tx.Rebind(expandedGenres)

	result, err := tx.ExecContext(ctx, expandedGenres, args...)

	if err != nil {
		return model.Movie{}, nil, err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return model.Movie{}, nil, err
	}

	if int64(len(genreIDs)) != rowAffected {
		return model.Movie{}, nil, ErrMovieGenreInsertMismatch
	}

	err = tx.Commit()

	if err != nil {
		return model.Movie{}, nil, err
	}

	return createdMovie, genres, nil
}
