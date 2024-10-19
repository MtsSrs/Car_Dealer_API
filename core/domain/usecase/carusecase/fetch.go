package carusecase

import (
	"github.com/mtssrs/car_dealer_API/core/domain"
	"github.com/mtssrs/car_dealer_API/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Car], error) {
	cars, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return cars, nil
}
