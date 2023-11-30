package gateway

import (
	"context"
	"encoding/json"
	"log"
	"microservices/libraryManager/internal/models"
	"net/http"
	"strconv"
)

type Gateway struct {
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) GetUser(ctx context.Context, id int) (*models.RecordModel, error) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:3000/get", nil)

	if err != nil {
		return nil, err
	}

	request = request.WithContext(ctx)

	values := request.URL.Query()
	values.Add("id", strconv.Itoa(id))
	request.URL.RawQuery = values.Encode()
	log.Println(request.URL)
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	//log.Println(response)
	var user models.Response

	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		log.Println(err)
		return nil, nil
	}
	log.Println(user)

	return nil, nil

}
