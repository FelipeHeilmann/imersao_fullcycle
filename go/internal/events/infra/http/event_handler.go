package http

import (
	"encoding/json"
	"net/http"

	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/application/usecase"
)

type EventHandler struct {
	listEvents *usecase.ListEvents
	getEvent   *usecase.GetEvent
	buyTickets *usecase.BuyTickets
	listSpots  *usecase.ListSpots
}

func NewEventsHandler(
	listEvents *usecase.ListEvents,
	getEvent *usecase.GetEvent,
	buyTickets *usecase.BuyTickets,
	listSpots *usecase.ListSpots,
) *EventHandler {
	return &EventHandler{
		listEvents: listEvents,
		getEvent:   getEvent,
		buyTickets: buyTickets,
		listSpots:  listSpots,
	}
}

func (handler *EventHandler) ListEvents(res http.ResponseWriter, req *http.Request) {
	output, err := handler.listEvents.Execute()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(output)

}

func (handler *EventHandler) GetEvent(res http.ResponseWriter, req *http.Request) {
	eventId := req.PathValue("eventId")
	input := usecase.InputGetEvent{
		EventId: eventId,
	}
	output, err := handler.getEvent.Execute(input)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(output)
}

func (handler *EventHandler) ListSpots(res http.ResponseWriter, req *http.Request) {
	eventId := req.PathValue("eventId")
	input := usecase.InputListSpots{
		EventId: eventId,
	}
	output, err := handler.listSpots.Execute(input)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(output)
}

func (handler *EventHandler) BuyTickets(res http.ResponseWriter, req *http.Request) {
	var input usecase.InputBuyTicket
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := handler.buyTickets.Execute(input)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(output)
}
