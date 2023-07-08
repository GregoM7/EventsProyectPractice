package event

import "github.com/GregoM7/EventsProyectPractice/internal/domain"

type Service interface {
	ReadAllEvents() ([]domain.Event, error)
	ReadAllEventsWithState() ([]domain.Event, error)
	Create(patient domain.Event) error
	Update(id int, patient domain.Event) error
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadAllEvents() ([]domain.Event, error) {
	l, err := s.r.ReadAllEvents()
	if err != nil {
		return nil, err
	}
	return l, nil
}
func (s *service) ReadAllEventsWithState() ([]domain.Event, error) {
	l, err := s.r.ReadAllEventsWithState()
	if err != nil {
		return nil, err
	}
	return l, nil
}
func (s *service) Create(event domain.Event) error {
	err := s.r.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) Update(id int,event domain.Event) error {
	err := s.r.UpdateEvent(id,event)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) Delete(id int) error {
	err := s.r.DeleteEvent(id)
	if err != nil {
		return err
	}
	return nil
}
