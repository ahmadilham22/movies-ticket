package model

import "time"

type Ticket struct {
	Id        string    `db:"id" json:"id"`
	EventName string    `db:"event_name" json:"event_name"`
	Price     int       `db:"price" json:"price"`
	Quota     int       `db:"quota" json:"quota"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type BuyTicketRequest struct {
	TicketID string `json:"ticket_id"`
	Quantity int    `json:"quantity"`
}
