package carservice

import (
	"encoding/json"
	"net/http"

	"github.com/mtssrs/car_dealer_API/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	carRequest, err := dto.FromJSONCreateCarRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	car, err := service.usecase.Create(carRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(car)
}
