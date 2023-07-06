package user

import (
	"errors"

	"github.com/GregoM7/EventsProyectPractice/internal/domain"
	"github.com/GregoM7/EventsProyectPractice/package/store"
)

type Repository interface {
	ReadAll() ([]domain.User, error)
	CreateUser(user domain.User) error
	ExistsUserByUsername(username string) bool
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

//CreateUser
func (r *repository) CreateUser(user domain.User) error {
	if r.ExistsUserByUsername(user.Username){
		return errors.New("Username already exists")
	}
	err := r.store.CreateUser(user)
	if err != nil {
		return errors.New("Error creating a User")
	}
	return nil
}

//UserCheck
func (r *repository) ExistsUserByUsername(username string) bool {
	 boolean := r.store.ExistsUserByUsername(username)
	 return boolean
}



