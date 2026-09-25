package model

import "time"

type CreateTicketRequest struct {
	MovieID  string    `json:"movie_id" binding:"required,uuid"`
	StartsAt time.Time `json:"starts_at" binding:"required"`
	Price    int       `json:"price" binding:"gt=0,lte=2147483647"`
	Quota    *int      `json:"quota" binding:"required,gte=0,lte=2147483647"`
}
