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
	Rating10    Rating = "10"
	Rating12    Rating = "12"
	Rating14    Rating = "14"
	Rating16    Rating = "16"
	Rating18    Rating = "18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}

	if e.Capacity <= 0 {
		return ErrEventCapacityZero
	}

	if e.Capacity <= 0 {
		return ErrEventPriceZero
	}

	return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)

	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)
	return spot, nil

}
