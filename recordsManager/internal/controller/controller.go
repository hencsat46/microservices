package controller

import (
	"microservices/recordsManager/internal/gateway"
	"microservices/recordsManager/internal/handler"
	"microservices/recordsManager/internal/models"
)

type usecase struct {
	gateway *gateway.Gateway
}

func NewUsecase(gateway *gateway.Gateway) handler.UsecaseInterfaces {
	return &usecase{gateway: gateway}
}

func (u *usecase) CreateUser(user models.User) error {
	return nil
}

func (u *usecase) ReadUser(id int) (*models.User, error) {
	return nil, nil
}

func (u *usecase) UpdateUser(user models.User) error {
	return nil
}

func (u *usecase) DeleteUser(id int) error {
	return nil
}

func (u *usecase) CreateRecord(user models.RecordModel) error {
	return nil
}

func (u *usecase) ReadRecord(id int) (*models.RecordModel, error) {
	return nil, nil
}

func (u *usecase) UpdateRecord(user models.RecordModel) error {
	return nil
}

func (u *usecase) DeleteRecord(id int) error {
	return nil
}
