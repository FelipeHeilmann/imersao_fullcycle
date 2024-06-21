package usecase

import "github.com/FelipeHeilmann/imersao_fullcycle/internal/events/domain"

type InputListSpots struct {
	EventId string `json:"event_id"`
}

type OutputListSpots struct {
	Event EventDto     `json:"event"`
	Spots []OutputSpot `json:"spots"`
}

type OutputSpot struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	EventId  string `json:"eventId"`
	Reserved bool   `json:"reserved"`
	Status   string `json:"status"`
	TicketId string `json:"ticketId"`
}

type ListSpots struct {
	repository domain.EventRepository
}

func NewListSpots(repository domain.EventRepository) *ListSpots {
	return &ListSpots{repository: repository}
}

func (listSpots *ListSpots) Execute(input InputListSpots) (*OutputListSpots, error) {
	event, err := listSpots.repository.FindById(input.EventId)
	if err != nil {
		return nil, err
	}

	spots, err := listSpots.repository.FindSpotsByEventId(input.EventId)
	if err != nil {
		return nil, err
	}

	outputSpots := make([]OutputSpot, len(spots))
	for i, spot := range spots {
		outputSpots[i] = OutputSpot{
			Id:       spot.Id,
			Name:     spot.Name,
			Status:   string(spot.Status),
			TicketId: spot.TicketId,
		}
	}

	outputEvent := EventDto{
		Id:           event.Id,
		Name:         event.Name,
		Location:     event.Location,
		Organization: event.Organization,
		Rating:       string(event.Rating),
		Date:         event.Date.Format("2006-01-02 15:04:05"),
		Capacity:     event.Capacity,
		Price:        event.Price,
		PartnerId:    event.PartnerId,
		ImageUrl:     event.ImageUrl,
	}

	return &OutputListSpots{Event: outputEvent, Spots: outputSpots}, nil
}
