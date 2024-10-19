package productrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/mtssrs/car_dealer_API/core/domain"
	"github.com/mtssrs/car_dealer_API/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Car], error) {
	ctx := context.Background()
	cars := []domain.Car{}
	total := int32(0)

	p := paginate.Instance(domain.Car{})

	query, queryCount := p.
		Query("SELECT * FROM cars").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "car_name", "car_manufacturer").
		Select()

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			car := domain.Car{}

			rows.Scan(
				&car.CarId,
				&car.CarManufacturer,
				&car.CarName,
				&car.CarModel,
				&car.CarYearModel,
				&car.CarPrice,
			)

			cars = append(cars, car)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Car]{
		Items: cars,
		Total: total,
	}, nil

}
