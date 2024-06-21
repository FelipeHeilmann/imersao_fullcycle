package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseUrl string
}

type Partner1ReservationRequest struct {
	Spot       []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	Email      string   `json:"email"`
}

type Partner1ReservationResponse struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventId    string `json:"event_id"`
}

func (parter *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	parterReq := Partner1ReservationRequest{
		Spot:       req.Spots,
		TicketKind: req.TicketType,
		Email:      req.Email,
	}

	body, err := json.Marshal(parterReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserve", parter.BaseUrl, req.EventId)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}
	var parterResponse []Partner1ReservationResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&parterResponse); err != nil {
		return nil, err
	}
	responses := make([]ReservationResponse, len(parterResponse))
	for i, r := range parterResponse {
		responses[i] = ReservationResponse{
			Id:     r.Id,
			Email:  r.Email,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}
	return responses, nil
}
