package repostitory

import (
	"microservices/authentication/internal/controller"
	"microservices/authentication/internal/models"
	"sync"
)

type repo struct {
	sync.RWMutex
	database map[int]mockUsers
}

func NewRepository() controller.RepositoryInterfaces {
	return &repo{database: make(map[int]mockUsers)}
}

func (r *repo) Create(user models.User, id int) error {
	r.Lock()
	defer r.Unlock()
	u := mockUsers{user.Name, user.Surname, user.Username, user.Password}
	r.database[id] = u
	return nil
}

func (r *repo) Read(user models.User, id int) (models.User, error) {
	r.RLock()
	defer r.RUnlock()
	value, ok := r.database[id]
	if !ok {
		return models.User{}, nil
	}

	return models.User{Name: value.Name, Surname: value.Surname, Username: value.Username, Password: ""}, nil
}

func (r *repo) Update(user models.User, id int) error {
	r.Lock()
	defer r.Unlock()
	value, ok := r.database[id]
	if !ok {
		return nil
	}
	r.database[id] = value
	return nil
}

func (r *repo) Delete(user models.User, id int) error {
	r.Lock()
	defer r.Unlock()
	_, ok := r.database[id]
	if !ok {
		return nil
	}
	delete(r.database, id)
	return nil
}
