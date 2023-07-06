package event

import "github.com/GregoM7/EventsProyectPractice/internal/domain"

type Service interface {
	ReadAllEvents() ([]domain.Event, error)
	ReadAllEventsWithState() ([]domain.Event, error)

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