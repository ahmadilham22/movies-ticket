package model

import "time"

type Ticket struct {
	Id        string    `db:"id" json:"id"`
	Price     int       `db:"price" json:"price"`
	Quota     int       `db:"quota" json:"quota"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	MovieId   string    `db:"movie_id" json:"movie_id"`
	StartsAt  time.Time `db:"starts_at" json:"starts_at"`
}

type BuyTicketRequest struct {
	TicketID string `json:"ticket_id"`
	Quantity int    `json:"quantity"`
}
