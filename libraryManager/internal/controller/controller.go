package controller

import (
	"context"
	"log"
	"microservices/libraryManager/internal/gateway"
	"microservices/libraryManager/internal/handler"
	"microservices/libraryManager/internal/models"
	"time"
)

type RepositoryInterfaces interface {
	Create(models.RecordModel, int) error
	Read(int) (models.RecordModel, error)
	Update(models.RecordModel, int) error
	Delete(int) error
}

type controller struct {
	repo    RepositoryInterfaces
	gateway gateway.Gateway
}

func NewUsecase(repo RepositoryInterfaces, gateway gateway.Gateway) handler.UsecaseInterfaces {
	return &controller{repo: repo, gateway: gateway}
}

func (c *controller) Create(user models.RecordModel, id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c.gateway.GetUser(ctx, id)

	// err := c.repo.Create(user, id)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// return nil
	return nil
}

func (c *controller) Read(id int) (models.RecordModel, error) {
	data, err := c.repo.Read(id)
	if err != nil {
		log.Println(err)
		return models.RecordModel{}, nil
	}
	return data, nil
}

func (c *controller) Update(user models.RecordModel, id int) error {
	err := c.repo.Update(user, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *controller) Delete(id int) error {
	err := c.repo.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
