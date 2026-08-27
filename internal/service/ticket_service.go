package service

import (
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
)

var ErrInvalidQuantity = errors.New("invalid quantity")
var ErrTicketSoldOut = errors.New("ticket sold out")

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
	result, err := t.tp.CreateTicket(ticket)
	if err != nil {
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
