package domain

import "errors"

var ErrTicketPriceZero = errors.New("ticket price must be greater than zero")

type TicketType = string

const (
	TicketTypeHalf TicketType = "half" //half=price ticket
	TicketTypeFull TicketType = "full" //full=price ticket
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func IsValidTycketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) ValidatePrice() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}
