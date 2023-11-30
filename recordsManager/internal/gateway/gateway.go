package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"microservices/recordsManager/internal/models"
	"net/http"
)

type Gateway struct {
}

func New() *Gateway {
	return &Gateway{}
}

func (g *Gateway) CreateUser(ctx context.Context, user models.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost:3000/create", bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	request = request.WithContext(ctx)
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	var responseModel models.Response

	if err := json.NewDecoder(response.Body).Decode(&responseModel); err != nil {
		return err
	}
	log.Println(response)
	log.Println(responseModel)

	return nil

}

func (g *Gateway) ReadUser(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

func (g *Gateway) UpdateUser(ctx context.Context, user models.User, id int) error {
	return nil
}

func (g *Gateway) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (g *Gateway) CreateRecord(ctx context.Context, user models.RecordModel, id int) error {
	return nil
}

func (g *Gateway) ReadRecord(ctx context.Context, id int) (*models.RecordModel, error) {
	return nil, nil
}

func (g *Gateway) UpdateRecord(ctx context.Context, user models.RecordModel, id int) error {
	return nil
}

func (g *Gateway) DeleteRecord(ctx context.Context, id int) error {
	return nil
}
