package carusecase

import (
	"github.com/mtssrs/car_dealer_API/core/domain"
	"github.com/mtssrs/car_dealer_API/core/dto"
)

func (usecase usecase) Create(carRequest *dto.CreateCarRequest) (*domain.Car, error) {
	car, err := usecase.repository.Create(carRequest)

	if err != nil {
		return nil, err
	}

	return car, nil
}
