package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketType string

var (
	ErrInvalidTicketType = errors.New("invalid ticket type")
)

const (
	TicketTypeFull TicketType = "full"
	TicketTypeHalf TicketType = "Half"
)

type Ticket struct {
	Id         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
	}

	ticket := &Ticket{
		Id:         uuid.New().String(),
		EventId:    event.Id,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) CalculatePrice() {
	if ticket.TicketType == TicketTypeHalf {
		ticket.Price = ticket.Price / 2
	}
}

func (ticket *Ticket) Validate() error {
	if ticket.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}
