package controller

import (
	"context"
	"log"
	"microservices/recordsManager/internal/gateway"
	"microservices/recordsManager/internal/handler"
	"microservices/recordsManager/internal/models"
	"time"
)

type usecase struct {
	gateway *gateway.Gateway
}

func NewUsecase(gateway *gateway.Gateway) handler.UsecaseInterfaces {
	return &usecase{gateway: gateway}
}

func (u *usecase) CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.gateway.CreateUser(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) ReadUser(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user, err := u.gateway.ReadUser(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (u *usecase) UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.gateway.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := u.gateway.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}

func (u *usecase) CreateRecord(record models.RecordModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.gateway.CreateRecord(ctx, record)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) ReadRecord(id int) (*models.RecordModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	record, err := u.gateway.ReadRecord(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return record, nil
}

func (u *usecase) UpdateRecord(record models.RecordModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.gateway.UpdateRecord(ctx, record)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *usecase) DeleteRecord(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := u.gateway.DeleteRecord(ctx, id); err != nil {
		return err
	}

	return nil
}
