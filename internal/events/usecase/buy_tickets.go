package usecase

import (
	"github.com/vinicius3g/golang/internal/events/domain"
	"github.com/vinicius3g/golang/internal/events/infra/service"
)

type BuyTicketsInputDto struct {
	EventID    string   `json:"event_id"`
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticket_kind"`
	CardHash   string   `json:"card_hash"`
	Email      string   `json:"email"`
}

type BuyTicketsOutputDto struct {
	Tickets []TicketDto `json:"tickets"`
}

type BuyTicketsUseCase struct {
	repo           domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTicketsUseCase(repo domain.EventRepository, partnerFactory service.PartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{repo: repo, partnerFactory: partnerFactory}
}

func (uc *BuyTicketsUseCase) Execute(input BuyTicketsInputDto) (*BuyTicketsOutputDto, error) {
	// Verifica o evento
	event, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}
	// Cria a solicitação de reserva
	req := &service.ReservationRequest{
		EventID:    input.EventID,
		Spots:      input.Spots,
		TicketKind: input.TicketKind,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}
	// Obtém o serviço do parceiro
	partnerService, err := uc.partnerFactory.CreatePartner(event.PartnerID)
	if err != nil {
		return nil, err
	}
	// Reserva os lugares usando o serviço do parceiro
	reservetionResponse, err := partnerService.MakeReservation(req)
	if err != nil {
		return nil, err
	}
	// Salva os ingressos no banco de dados
	tickets := make([]domain.Ticket, len(reservetionResponse))
	for i, reservation := range reservetionResponse {
		spot, err := uc.repo.FindSpotByName(event.ID, reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket, err := domain.NewTicket(event, spot, domain.TicketKind(reservation.TicketKind))
		if err != nil {
			return nil, err
		}

		err = uc.repo.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.Reserve(ticket.ID)
		err = uc.repo.ReserveSpot(spot.ID, ticket.ID)
		if err != nil {
			return nil, err
		}

		tickets[i] = *ticket
	}

	ticketsDtos := make([]TicketDto, len(tickets))
	for i, ticket := range tickets {
		ticketsDtos[i] = TicketDto{
			ID:         ticket.ID,
			SpotID:     ticket.Spot.ID,
			TicketKind: string(ticket.TicketKind),
			Price:      ticket.Price,
		}
	}

	return &BuyTicketsOutputDto{Tickets: ticketsDtos}, nil

}
