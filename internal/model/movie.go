package model

import "time"

type Movie struct {
	ID              string    `db:"id"`
	Title           string    `db:"title"`
	Synopsis        string    `db:"synopsis"`
	DurationMinutes int       `db:"duration_minutes"`
	ReleaseDate     time.Time `db:"release_date"`
	PosterURL       string    `db:"poster_url"`
	TrailerURL      *string   `db:"trailer_url"`
	AgeRating       string    `db:"age_rating"`
	Language        string    `db:"language"`
	Country         string    `db:"country"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type MovieWithGenres struct {
	Movie  Movie
	Genres []Genre
}
