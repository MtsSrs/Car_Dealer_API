package carusecase

import "github.com/mtssrs/car_dealer_API/core/domain"

type usecase struct {
	repository domain.CarRepository
}

func New(repository domain.CarRepository) domain.CarUseCase {
	return &usecase{
		repository: repository,
	}
}
