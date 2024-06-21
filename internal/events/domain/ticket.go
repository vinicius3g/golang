package domain

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
