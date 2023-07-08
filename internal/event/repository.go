package event

import (
	"errors"

	"github.com/GregoM7/EventsProyectPractice/internal/domain"
	"github.com/GregoM7/EventsProyectPractice/package/store"
)

type Repository interface {
	ReadAllEvents() ([]domain.Event, error)
	ReadAllEventsWithState() ([]domain.Event, error)
	CreateEvent(event domain.Event) error
	DeleteEvent(id int) error
	UpdateEvent(id int, event domain.Event) error
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


func (r *repository) CreateEvent(event domain.Event) error{
	err := r.store.CreateEvent(event)
	if err != nil {
		return errors.New("Error creating a new Event")
	}
	return nil
}
func (r *repository) DeleteEvent(id int) error {
	err := r.store.DeleteEvent(id)
	if err != nil {
		return errors.New("Error deleting a Event - Cause 1:he have still inscriptions. Cause 2: He doest exist.")
	}
	return nil
}
func (r *repository) UpdateEvent(id int, event domain.Event) error {
	original, err := r.store.ReadEventById(id)
	if err != nil {
		return errors.New("The Event does not exists")
	}
	complete := unchangeEmptysEvent(event, original)
	err = r.store.UpdateEvent(id, complete)
	if err != nil {
		return errors.New("Error updating a Event")
	}
	return nil
}

// Completando el update con Original
func unchangeEmptysEvent(event domain.Event, original domain.Event) domain.Event {

	if event.Titulo == "" {
		event.Titulo = original.ShortDescription
	}
	if event.ShortDescription == "" {
		event.ShortDescription = original.ShortDescription
	}
	if event.LongDescription == "" {
		event.LongDescription = original.LongDescription
	}
	if event.State == "" {
		event.State = original.State
	}
	if event.FechaYHora.GoString() == "" {
		event.FechaYHora = original.FechaYHora
	}

	return event
}