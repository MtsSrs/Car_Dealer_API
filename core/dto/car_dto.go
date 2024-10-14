package dto

import (
	"encoding/json"
	"io"
)

type CreateCarRequest struct {
	CarManufacturer string  `json:"carManufacturer"`
	CarName         string  `json:"carName"`
	CarModel        string  `json:"carModel"`
	CarYearModel    string  `json:"carYearModel"`
	CarPrice        float32 `json:"carPrice"`
}

func FromJSONCreateCarRequest(body io.Reader) (*CreateCarRequest, error) {
	createCarRequest := CreateCarRequest{}
	if err := json.NewDecoder(body).Decode(&createCarRequest); err != nil {
		return nil, err
	}

	return &createCarRequest, nil
}
