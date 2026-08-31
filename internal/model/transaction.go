package model

import "time"

type Transaction struct {
	Id          string    `db:"id" json:"id"`
	TicketID    string    `db:"ticket_id" json:"ticket_id"`
	UserID      string    `db:"user_id" json:"user_id"`
	BookingCode string    `db:"booking_code" json:"booking_code"`
	Quantity    int       `db:"quantity" json:"quantity"`
	TotalPrice  int       `db:"total_price" json:"total_price"`
	Status      string    `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
