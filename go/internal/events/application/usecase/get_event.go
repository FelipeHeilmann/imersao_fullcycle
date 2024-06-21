package usecase

import (
	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/domain"
)

type InputGetEvent struct {
	EventId string
}

type OutputGetEvent struct {
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

type GetEvent struct {
	repository domain.EventRepository
}

func NewGetEvent(repository domain.EventRepository) *GetEvent {
	return &GetEvent{repository: repository}
}

func (getEvent *GetEvent) Execute(input InputGetEvent) (*OutputGetEvent, error) {
	event, err := getEvent.repository.FindById(input.EventId)
	if err != nil {
		return nil, err
	}
	return &OutputGetEvent{
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
	}, nil
}
