package usecase

import (
	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/domain"
	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/infra/service"
)

type InputBuyTicket struct {
	EventId    string   `json:"eventId"`
	Spots      []string `json:"spots"`
	TicketType string   `json:"ticketType"`
	CardHash   string   `json:"cardHash"`
	Email      string   `json:"email"`
}

type OutputBuyTicket struct {
	Tickets []OutputTicket `json:"tickets"`
}

type OutputTicket struct {
	Id         string  `json:"id"`
	SpotId     string  `json:"spotId"`
	TicketType string  `json:"ticketType"`
	Price      float64 `json:"price"`
}

type BuyTickets struct {
	repository     domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTickets(repository domain.EventRepository, factory service.PartnerFactory) *BuyTickets {
	return &BuyTickets{repository: repository, partnerFactory: factory}
}

func (buyTicket *BuyTickets) Execute(input InputBuyTicket) (*OutputBuyTicket, error) {
	// Verifica o evento
	event, err := buyTicket.repository.FindById(input.EventId)
	if err != nil {
		return nil, err
	}

	// Cria a solicitação de reserva
	req := &service.ReservationRequest{
		EventId:    input.EventId,
		Spots:      input.Spots,
		TicketType: input.TicketType,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}

	// Obtém o serviço do parceiro
	partnerService, err := buyTicket.partnerFactory.Create(event.PartnerId)
	if err != nil {
		return nil, err
	}

	// Reserva os lugares usando o serviço do parceiro
	reservationResponse, err := partnerService.MakeReservation(req)
	if err != nil {
		return nil, err
	}

	tickets := make([]domain.Ticket, len(reservationResponse))
	for i, reservation := range reservationResponse {
		spot, err := buyTicket.repository.FindSpotByName(event.Id, reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket, err := domain.NewTicket(event, spot, domain.TicketType(input.TicketType))
		if err != nil {
			return nil, err
		}

		err = buyTicket.repository.SaveTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.Reserve(ticket.Id)
		err = buyTicket.repository.ReserveSpot(spot.Id, ticket.Id)
		if err != nil {
			return nil, err
		}

		tickets[i] = *ticket
	}

	output := make([]OutputTicket, len(tickets))
	for i, ticket := range tickets {
		output[i] = OutputTicket{
			Id:         ticket.Id,
			SpotId:     ticket.Spot.Id,
			TicketType: string(ticket.TicketType),
			Price:      ticket.Price,
		}
	}

	return &OutputBuyTicket{Tickets: output}, nil
}
