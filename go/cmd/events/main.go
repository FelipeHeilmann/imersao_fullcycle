package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/application/usecase"
	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/infra/repository"
	"github.com/FelipeHeilmann/imersao_fullcycle/internal/events/infra/service"

	httpHandler "github.com/FelipeHeilmann/imersao_fullcycle/internal/events/infra/http"
)

func main() {
	db, err := sql.Open("postgres", "docker:123456@tcp(go-postgres:5431)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	eventRepository, err := repository.NewPostgresEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:3000/partner1",
		2: "http://localhost:3001/partner2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEvents := usecase.NewListEvents(eventRepository)
	getEvent := usecase.NewGetEvent(eventRepository)
	buyTickets := usecase.NewBuyTickets(eventRepository, partnerFactory)
	listSpots := usecase.NewListSpots(eventRepository)

	eventsHandler := httpHandler.NewEventsHandler(
		listEvents,
		getEvent,
		buyTickets,
		listSpots,
	)

	router := http.NewServeMux()
	router.HandleFunc("/events", eventsHandler.ListEvents)
	router.HandleFunc("/events/{eventId}", eventsHandler.GetEvent)
	router.HandleFunc("/events/{eventId}/spots", eventsHandler.ListSpots)
	router.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	http.ListenAndServe(":8080", router)

}
