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

	query, queryCount, err := paginate.Pagination("SELECT * FROM cars").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "car_name", "car_manufacturer").
		Query()

}
