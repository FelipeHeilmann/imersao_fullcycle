package usecase

import "github.com/FelipeHeilmann/imersao_fullcycle/internal/events/domain"

type OutputListEvents struct {
	Events []EventDto
}

type EventDto struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Organization string  `json:"organization"`
	Rating       string  `json:"rating"`
	Date         string  `json:"date"`
	Capacity     int     `json:"capacity"`
	Price        float64 `json:"price"`
	PartnerId    int     `json:"partnerId"`
	ImageUrl     string  `json:"imageUrl"`
}

type ListEvents struct {
	repository domain.EventRepository
}

func NewListEvents(repository domain.EventRepository) *ListEvents {
	return &ListEvents{repository: repository}
}

func (listEvent *ListEvents) Execute() (*OutputListEvents, error) {
	events, err := listEvent.repository.List()
	if err != nil {
		return nil, err
	}

	output := make([]EventDto, len(events))
	for i, event := range events {
		output[i] = EventDto{
			Id:           event.Id,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date.Format("2006-01-02 15:04:05"),
			ImageUrl:     event.ImageUrl,
			Capacity:     event.Capacity,
			Price:        event.Price,
			PartnerId:    event.PartnerId,
		}
	}
	return &OutputListEvents{Events: output}, nil
}
