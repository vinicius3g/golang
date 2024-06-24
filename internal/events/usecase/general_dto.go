package usecase

type EventDto struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	ImageURL     string  `json:"image_url"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerID    int     `json:"partner_id"`
}

type SpotDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Reserved bool   `json:"reserved"`
	TicketID string `json:"ticket_id"`
}
