package domain

import (
	"errors"

	"github.com/google/uuid"
)

type SpotStatus string

var (
	ErrInvalidSpotNumber   = errors.New("invalid spot number")
	ErrSpotNotFound        = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
)

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	Id       string
	EventId  string
	Name     string
	Status   SpotStatus
	TicketId string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		Id:      uuid.New().String(),
		EventId: event.Id,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	v := spot.Validade()
	if v != nil {
		return nil, v
	}

	return spot, nil
}

func (spot *Spot) Validade() error {
	if len(spot.Name) == 0 {
		return errors.New("spot name is required")
	}

	if len(spot.Name) < 2 {
		return errors.New("spot name must be as least 2 characteres long")
	}

	if spot.Name[0] < 'A' || spot.Name[0] > 'Z' {
		return errors.New("spot name must start with letters")
	}

	if spot.Name[1] < '0' || spot.Name[1] > '9' {
		return errors.New("spot name must end with numbers")
	}

	return nil
}

func (spot *Spot) Reserve(ticketId string) error {
	if spot.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}

	spot.Status = SpotStatusSold
	spot.TicketId = ticketId
	return nil
}
