package model

import "time"

type CreateMovieRequest struct {
	Title           string   `json:"title" binding:"required"`
	Synopsis        string   `json:"synopsis" binding:"required"`
	DurationMinutes int      `json:"duration_minutes" binding:"required"`
	ReleaseDate     string   `json:"release_date" binding:"required"`
	PosterURL       string   `json:"poster_url" binding:"required"`
	TrailerURL      *string  `json:"trailer_url"`
	AgeRating       string   `json:"age_rating" binding:"required"`
	Language        string   `json:"language" binding:"required"`
	Country         string   `json:"country" binding:"required"`
	GenreIDs        []string `json:"genre_ids" binding:"required,min=1,dive,uuid"`
}

type GenreSummaryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MovieResponse struct {
	ID              string                 `json:"id"`
	Title           string                 `json:"title"`
	Synopsis        string                 `json:"synopsis"`
	DurationMinutes int                    `json:"duration_minutes"`
	ReleaseDate     string                 `json:"release_date"`
	PosterURL       string                 `json:"poster_url"`
	TrailerURL      *string                `json:"trailer_url"`
	AgeRating       string                 `json:"age_rating"`
	Language        string                 `json:"language"`
	Country         string                 `json:"country"`
	Genres          []GenreSummaryResponse `json:"genres"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}
