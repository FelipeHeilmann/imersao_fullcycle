package domain

import "errors"

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

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) CalculatePrice() {
	if ticket.TicketType == TicketTypeHalf {
		ticket.Price = ticket.Price / 2
	}
}

func (ticket *Ticket) Validade() error {
	if ticket.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}
