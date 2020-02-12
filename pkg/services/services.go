package services

import "encoding/json"

type Services []struct {
	Name       string       `json:"name"`
	Objectives []Objectives `json:"objectives"`
}

func NewServicesFromJSON(data []byte) (*Services, error) {
	var services Services
	err := json.Unmarshal(data, &services)
	if err != nil {
		return nil, err
	}

	return &services, nil
}
