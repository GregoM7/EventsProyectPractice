package user

import (
	"github.com/GregoM7/EventsProyectPractice/internal/domain"
	"github.com/GregoM7/EventsProyectPractice/package/store"
	
)

type Repository interface {
	ReadAll() ([]domain.User, error)
}

type repository struct {
	store store.Store
}

func NewRepository(store store.Store) Repository {
	return &repository{store: store}
}

// ReadAll 
func (r *repository) ReadAll() ([]domain.User, error) {
	list, err := r.store.ReadAllUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return list, nil
}


