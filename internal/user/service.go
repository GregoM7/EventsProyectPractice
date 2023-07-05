package user

import "github.com/GregoM7/EventsProyectPractice/internal/domain"

type Service interface {
	ReadAll() ([]domain.User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadAll() ([]domain.User, error){
	l, err := s.r.ReadAll()
	if err != nil {
		return nil, err
	}
	return l, nil
}