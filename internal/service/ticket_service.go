package service

import (
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
	"time"
)

var ErrInvalidQuantity = errors.New("invalid quantity")
var ErrTicketSoldOut = errors.New("ticket sold out")
var ErrShowTimeNotInFuture = errors.New("showtime must be in the future")

type TicketService struct {
	tp *repository.TicketRepository
}

func NewTicketService(tp *repository.TicketRepository) *TicketService {
	return &TicketService{
		tp: tp,
	}
}

func (t *TicketService) FetchAllTickets() ([]model.Ticket, error) {
	result, err := t.tp.GetAllTickets()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TicketService) CreateTicket(ticket model.Ticket) (model.Ticket, error) {
	if !ticket.StartsAt.After(time.Now()) {
		return ticket, ErrShowTimeNotInFuture
	}

	result, err := t.tp.CreateTicket(ticket)
	if err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return ticket, ErrMovieNotFound
		}
		return ticket, err
	}

	return result, nil
}

func (t *TicketService) BuyTicket(ticket model.BuyTicketRequest, userId string) error {

	if ticket.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	err := t.tp.CreateBuy(ticket, userId)
	if err != nil {
		if errors.Is(err, repository.ErrTicketSoldOut) {
			return ErrTicketSoldOut
		}
		return err
	}

	return nil
}
