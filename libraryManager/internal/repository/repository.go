package repository

import (
	"microservices/libraryManager/internal/controller"
	"microservices/libraryManager/internal/models"
	"sync"
)

type repo struct {
	sync.RWMutex
	database map[int]mockRecords
}

func NewRepository() controller.RepositoryInterfaces {
	return &repo{database: make(map[int]mockRecords)}
}

func (r *repo) Create(user models.RecordModel, id int) error {
	r.Lock()
	defer r.Unlock()
	u := mockRecords{UserId: user.UserId}
	r.database[id] = u
	return nil
}

func (r *repo) Read(id int) (models.RecordModel, error) {
	r.RLock()
	defer r.RUnlock()
	value, ok := r.database[id]
	if !ok {
		return models.RecordModel{}, nil
	}

	return models.RecordModel{UserId: value.UserId, RecordId: id}, nil
}

func (r *repo) Update(user models.RecordModel, id int) error {
	r.Lock()
	defer r.Unlock()
	_, ok := r.database[id]
	if !ok {
		return nil
	}
	r.database[id] = mockRecords{UserId: user.UserId}
	return nil
}

func (r *repo) Delete(id int) error {
	r.Lock()
	defer r.Unlock()
	_, ok := r.database[id]
	if !ok {
		return nil
	}
	delete(r.database, id)
	return nil
}
