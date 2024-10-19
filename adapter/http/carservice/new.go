package carservice

import "github.com/mtssrs/car_dealer_API/core/domain"

type service struct {
	usecase domain.CarUseCase
}

func New(usecase domain.CarUseCase) domain.CarService {
	return &service{
		usecase: usecase,
	}
}
