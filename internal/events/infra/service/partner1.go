package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseURL string
}

type Partner1ReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticket_kind"`
	Email      string   `json:"email"`
}

type Partner1ReservationResponse struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventID    string `json:"event_id"`
}

func (p *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerReq := Partner1ReservationRequest{
		Spots:      req.Spots,
		TicketKind: req.TicketType,
		Email:      req.Email,
	}

	body, err := json.Marshal(partnerReq)

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserve", p.BaseURL, req.EventID)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	var partnerRes []Partner1ReservationResponse

	if err := json.NewDecoder(httpResp.Body).Decode(&partnerRes); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerRes))

	for i, r := range partnerRes {
		responses[i] = ReservationResponse{
			ID:     r.Id,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}

	return responses, nil
}
