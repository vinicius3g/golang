package usecase

import "github.com/vinicius3g/golang/internal/events/domain"

type ListSpotsInputDto struct {
	EventID string `json:"event_id"`
}

type ListSpotsOutputDto struct {
	Event EventDto  `json:"event"`
	Spots []SpotDto `json:"spots"`
}

type ListSpotsUseCase struct {
	repo domain.EventRepository
}

func NewListSpotsUseCase(repo domain.EventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{repo: repo}
}

func (uc *ListSpotsUseCase) Execute(input ListSpotsInputDto) (*ListSpotsOutputDto, error) {
	event, err := uc.repo.FindEventByID(input.EventID)
	if err != nil {
		return nil, err
	}

	spots, err := uc.repo.FindSpotsByEventID(input.EventID)
	if err != nil {
		return nil, err
	}

	spotsDto := make([]SpotDto, len(spots))
	for i, spot := range spots {
		spotsDto[i] = SpotDto{
			ID:       spot.ID,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketID: spot.TicketID,
		}
	}

	eventDto := EventDto{
		ID:           event.ID,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		ImageURL:     event.ImageURL,
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerID:    event.PartnerID,
	}

	return &ListSpotsOutputDto{Event: eventDto, Spots: spotsDto}, nil
}
