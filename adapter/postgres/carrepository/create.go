package carrepository

import (
	"context"

	"github.com/mtssrs/car_dealer_API/core/domain"
	"github.com/mtssrs/car_dealer_API/core/dto"
)

func (repository repository) Create(carRequest *dto.CreateCarRequest) (*domain.Car, error) {
	ctx := context.Background()
	car := domain.Car{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO cars (car_manufacturer, car_name, car_model, car_year_model, car_price) VALUES ($1, $2, $3, $4, $5) RETURNING *",
		carRequest.CarManufacturer,
		carRequest.CarName,
		carRequest.CarModel,
		carRequest.CarYearModel,
		carRequest.CarPrice,
	).Scan(
		&car.CarId,
		&car.CarManufacturer,
		&car.CarName,
		&car.CarModel,
		&car.CarYearModel,
		&car.CarPrice,
	)

	if err != nil {
		return nil, err
	}

	return &car, nil
}
