package domain

import (
	"errors"
)

// Errors
var (
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

// TicketKind represents the kind of a ticket.
type TicketKind string

const (
	TicketKindHalf TicketKind = "half" // Half-price ticket
	TicketKindFull TicketKind = "full" // Full-price ticket
)

// IsValidTicketKind checks if a ticket kind is valid.
func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

// Ticket represents a ticket for an event.
type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

// CalculatePrice calculates the price based on the ticket kind.
func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}

// Validate checks if the ticket data is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}
