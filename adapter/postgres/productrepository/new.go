package productrepository

import (
	"github.com/mtssrs/car_dealer_API/adapter/postgres"
	"github.com/mtssrs/car_dealer_API/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.CarRepository {
	return &repository{
		db: db,
	}
}
