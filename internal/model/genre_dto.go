package model

import "time"

type CreateGenreRequest struct {
	Name string `json:"name" binding:"required"`
}

type GenreResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
