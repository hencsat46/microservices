package controller

import (
	"log"
	"microservices/authentication/internal/handler"
	"microservices/authentication/internal/models"
)

type controller struct {
	repo RepositoryInterfaces
}

type RepositoryInterfaces interface {
	Create(models.User, int) error
	Read(int) (models.User, error)
	Update(models.User, int) error
	Delete(int) error
}

func NewUsecase(repo RepositoryInterfaces) handler.UsecaseInterfaces {
	return &controller{repo: repo}
}

func (c *controller) Create(user models.User, id int) error {
	err := c.repo.Create(user, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *controller) Read(id int) (models.User, error) {
	data, err := c.repo.Read(id)
	if err != nil {
		log.Println(err)
		return models.User{}, nil
	}
	return data, nil
}

func (c *controller) Update(user models.User, id int) error {
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
