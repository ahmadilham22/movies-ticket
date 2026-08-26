package repository

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"online-ticketing/internal/model"

	"github.com/jmoiron/sqlx"
)

var ErrTicketSoldOut = errors.New("ticket sold out")

type TicketRepository struct {
	db *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (t *TicketRepository) GetAllTickets() ([]model.Ticket, error) {
	tickets := []model.Ticket{}
	err := t.db.Select(&tickets, "SELECT * FROM tickets ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (t *TicketRepository) CreateTicket(ticket model.Ticket) (model.Ticket, error) {
	var result model.Ticket
	tx, err := t.db.Beginx()
	if err != nil {
		return ticket, err
	}

	defer tx.Rollback()

	query := "INSERT INTO tickets (event_name, price, quota) VALUES ($1, $2, $3) RETURNING *"

	err = tx.Get(&result, query, ticket.EventName, ticket.Price, ticket.Quota)

	if err != nil {
		return ticket, err
	}
	tx.Commit()
	return result, nil
}

func (t *TicketRepository) CreateBuy(req model.BuyTicketRequest, userId string) error {
	var ticketData model.Ticket

	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.Get(&ticketData, "SELECT price, quota FROM tickets WHERE id = $1 FOR UPDATE", req.TicketID)

	if err != nil {
		return err
	}

	if ticketData.Quota <= 0 {
		return ErrTicketSoldOut
	}

	if ticketData.Quota < req.Quantity {
		return ErrTicketSoldOut
	}

	totalPrice := ticketData.Price * req.Quantity
	status := "SUCCESS"
	randomByte := make([]byte, 8)
	if _, err := rand.Read(randomByte); err != nil {
		return err
	}

	bookingCode := hex.EncodeToString(randomByte)

	_, err = tx.Exec("UPDATE tickets SET quota = quota - $1 WHERE id = $2", req.Quantity, req.TicketID)

	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (ticket_id, user_id, booking_code, quantity, total_price, status) VALUES ($1, $2, $3, $4, $5, $6)", req.TicketID, userId, bookingCode, req.Quantity, totalPrice, status)

	if err != nil {
		return err
	}

	return tx.Commit()
}
