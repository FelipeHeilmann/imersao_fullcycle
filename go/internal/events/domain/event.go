package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventDateFuture   = errors.New("event date must be in the future")
	ErrEventCapacityZero = errors.New("event capacity must be greater than zero")
	ErrEventPriceZero    = errors.New("event price must be greater than zero")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	Id           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageUrl     string
	Capacity     int
	Price        float64
	PartnerId    int
	Spots        []Spot
	Tickets      []Ticket
}

func (event *Event) Validate() error {
	if event.Name == "" {
		return ErrEventNameRequired
	}

	if event.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}

	if event.Capacity <= 0 {
		return ErrEventCapacityZero
	}

	if event.Price <= 0 {
		return ErrEventPriceZero
	}
	return nil
}

func (event *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(event, name)
	if err != nil {
		return nil, err
	}
	event.Spots = append(event.Spots, *spot)
	return spot, nil
}
