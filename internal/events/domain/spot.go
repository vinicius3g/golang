package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("Spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErroSpotNameRequired       = errors.New("spot name is required")
	ErrSpotNameTwoCaracters    = errors.New("spot name must be at least with characters long")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndsWithNumber  = errors.New("spot name must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	// maneira 1 de retornar a validação
	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil

	// maneira 2 de retornar a validação

	// v := spot.Validate()

	// if v != nil {
	// 	return nil, v
	// }

	// return spot, nil
}

func (s Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErroSpotNameRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameTwoCaracters
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}

	if s.Name[1] < 0 || s.Name[1] > 9 {
		return ErrSpotNameEndsWithNumber
	}

	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}

	s.Status = SpotStatusSold
	s.TicketID = ticketID

	return nil
}
