package domain

type EventRepository interface {
	List() ([]Event, error)
	FindById(eventID string) (*Event, error)
	FindSpotsByEventId(eventID string) ([]*Spot, error)
	FindSpotByName(eventId, spotName string) (*Spot, error)
	Save(event *Event) error
	SaveSpot(spot *Spot) error
	SaveTicket(ticket *Ticket) error
	ReserveSpot(spotId, ticketId string) error
}
