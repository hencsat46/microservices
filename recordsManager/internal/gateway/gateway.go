package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"microservices/recordsManager/internal/models"
	"net/http"
	"strconv"
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

	var responseModel models.Response

	if err := json.NewDecoder(response.Body).Decode(&responseModel); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(responseModel)

	typeMap := responseModel.Payload.(map[string]interface{})

	return &models.User{Name: typeMap["Name"].(string), Surname: typeMap["Surname"].(string), Username: typeMap["Username"].(string), Password: "", Id: int(typeMap["Id"].(float64))}, nil
}

func (g *Gateway) UpdateUser(ctx context.Context, user models.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, "http://localhost:3000/update", bytes.NewBuffer(body))

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

func (g *Gateway) DeleteUser(ctx context.Context, id int) error {
	body, err := json.Marshal(models.User{Id: id})
	if err != nil {
		log.Println(err)
		return err
	}

	request, err := http.NewRequest(http.MethodDelete, "http://localhost:3000/delete", bytes.NewBuffer(body))

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

func (g *Gateway) CreateRecord(ctx context.Context, record models.RecordModel) error {
	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost:3001/create", bytes.NewBuffer(body))

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

func (g *Gateway) ReadRecord(ctx context.Context, id int) (*models.RecordModel, error) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:3001/get", nil)

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

	var responseModel models.Response

	if err := json.NewDecoder(response.Body).Decode(&responseModel); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(responseModel)

	typeMap := responseModel.Payload.(map[string]interface{})

	return &models.RecordModel{UserId: int(typeMap["UserId"].(float64)), RecordId: int(typeMap["RecordId"].(float64))}, nil
}

func (g *Gateway) UpdateRecord(ctx context.Context, record models.RecordModel) error {
	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, "http://localhost:3001/update", bytes.NewBuffer(body))

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

func (g *Gateway) DeleteRecord(ctx context.Context, id int) error {
	body, err := json.Marshal(models.RecordModel{UserId: id})
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodDelete, "http://localhost:3001/delete", bytes.NewBuffer(body))

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
