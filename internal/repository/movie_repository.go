package repository

import (
	"context"
	"errors"
	"online-ticketing/internal/model"

	"github.com/jmoiron/sqlx"
)

type movieGenreRow struct {
	model.Movie
	GenreID   *string `db:"genre_id"`
	GenreName *string `db:"genre_name"`
}

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

func (r *MovieRepository) GetMovies(ctx context.Context) ([]model.MovieWithGenres, error) {
	rows := make([]movieGenreRow, 0)

	query := "SELECT m.id AS id, title, synopsis, duration_minutes, release_date, poster_url, trailer_url, age_rating, language, country, m.created_at, m.updated_at, g.id AS genre_id, g.name AS genre_name FROM movies m LEFT JOIN movie_genres mg ON mg.movie_id = m.id LEFT JOIN genres g ON g.id = mg.genre_id ORDER BY m.release_date DESC, m.title ASC, g.name ASC"

	err := r.db.SelectContext(ctx, &rows, query)

	if err != nil {
		return nil, err
	}

	results := make([]model.MovieWithGenres, 0)
	movieIndexes := make(map[string]int)
	for _, row := range rows {
		index, exists := movieIndexes[row.ID]
		if !exists {
			index = len(results)
			results = append(results, model.MovieWithGenres{
				Movie:  row.Movie,
				Genres: make([]model.Genre, 0),
			})
			movieIndexes[row.ID] = index
		}

		if row.GenreID != nil && row.GenreName != nil {
			results[index].Genres = append(results[index].Genres, model.Genre{ID: *row.GenreID, Name: *row.GenreName})
		}

	}

	return results, nil
}
