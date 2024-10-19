package di

import (
	"github.com/mtssrs/car_dealer_API/adapter/http/carservice"
	"github.com/mtssrs/car_dealer_API/adapter/postgres"
	"github.com/mtssrs/car_dealer_API/adapter/postgres/carrepository"
	"github.com/mtssrs/car_dealer_API/core/domain"
	"github.com/mtssrs/car_dealer_API/core/domain/usecase/carusecase"
)

func ConfigCarDi(conn postgres.PoolInterface) domain.CarService {
	carRepository := carrepository.New(conn)
	carUseCase := carusecase.New(carRepository)
	carService := carservice.New(carUseCase)

	return carService
}
