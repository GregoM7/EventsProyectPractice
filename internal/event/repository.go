package event

import (
	"github.com/GregoM7/EventsProyectPractice/internal/domain"
	"github.com/GregoM7/EventsProyectPractice/package/store"
)

type Repository interface {
	ReadAllEvents() ([]domain.Event, error)
	ReadAllEventsWithState() ([]domain.Event, error)
}

type repository struct {
	store store.Store
}

func NewRepository(store store.Store) Repository {
	return &repository{store: store}
}

func (r *repository) ReadAllEvents() ([]domain.Event, error) {
	list, err := r.store.ReadAllEvents()
	if err != nil {
		return []domain.Event{}, err
	}
	return list, nil
}
func (r *repository) ReadAllEventsWithState() ([]domain.Event, error) {
	list, err := r.store.ReadAllEventsWithState()
	if err != nil {
		return []domain.Event{}, err
	}
	return list, nil
}
