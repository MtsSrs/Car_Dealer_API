package carrepository

import (
	"context"
	"fmt"
	"strings"

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
		SearchBy(pagination.Search, "car_name").
		GroupBy("car_id").
		Select()

	// // Corrigir a consulta principal
	*query = strings.ReplaceAll(*query, "::TEXT", "")
	*query = strings.ReplaceAll(*query, "((", "((car_name")
	*query = strings.ReplaceAll(*query, ") or (", ") or (car_manufacturer")
	*query = strings.ReplaceAll(*query, "GROUP BY car_id", "")

	// // Corrigir a consulta de contagem
	*queryCount = strings.ReplaceAll(*queryCount, "COUNT(id)", "COUNT(*)")
	*queryCount = strings.ReplaceAll(*queryCount, "::TEXT", "")
	*queryCount = strings.ReplaceAll(*queryCount, "((", "((car_name")
	*queryCount = strings.ReplaceAll(*queryCount, ") or (", ") or (car_manufacturer")
	*queryCount = strings.ReplaceAll(*queryCount, "GROUP BY car_id", "")

	fmt.Println("Query corrigida:", *query)
	fmt.Println("Count Query corrigida:", *queryCount)
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
