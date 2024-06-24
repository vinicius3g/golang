package usecase

import "github.com/vinicius3g/golang/internal/events/domain"

type ListeventsOutputDto struct {
	Events []EventDto `json:"events"`
}

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventsUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}

func (uc *ListEventsUseCase) Execute() (*ListeventsOutputDto, error) {
	events, err := uc.repo.ListEvents()
	if err != nil {
		return nil, err
	}

	eventsDto := make([]EventDto, len(events))

	for i, event := range events {
		eventsDto[i] = EventDto{
			ID:           event.ID,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date.Format("2006-01 15:04:05"),
			ImageURL:     event.ImageURL,
			Capacity:     event.Capacity,
			Price:        event.Price,
			PartnerID:    event.PartnerID,
		}
	}

	return &ListeventsOutputDto{Events: eventsDto}, nil
}
