package domain

import (
	"net/http"

	"github.com/mtssrs/car_dealer_API/core/dto"
)

type Car struct {
	CarId        int32   `json:"carId"`
	CarName      string  `json:"carName"`
	CarModel     string  `json:"carModel"`
	CarYearModel string  `json:"carYearModel"`
	CarPrice     float32 `json:"carPrice"`
}

type CarService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

type CarUseCase interface {
	Create(carRequest *dto.CreateCarRequest) (*Car, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Car], error)
}

type CarRepository interface {
	Create(carRequest *dto.CreateCarRequest) (*Car, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Car], error)
}
