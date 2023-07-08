package store

import (
	"database/sql"
	"errors"

	"github.com/GregoM7/EventsProyectPractice/internal/domain"
	"github.com/GregoM7/EventsProyectPractice/internal/domain/dto"
)

type store struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) Store {
	return &store{db: db}
}

type Store interface {
	//------ USER
	ReadAllUsers() ([]domain.User, error)
	CreateUser(user domain.User) error
	ExistsUserByUsername(username string) bool
	GetUser(username string) (dto.UserGet, error)
	//------ EVENT
	ReadAllEvents() ([]domain.Event, error)
	ReadAllEventsWithState() ([]domain.Event, error)
	ReadEventById(id int) (domain.Event, error)
	CreateEvent(event domain.Event) error
	DeleteEvent(id int) error
	UpdateEvent(id int, event domain.Event) error
	//------ INSCRIPTION
}

func (s *store) ReadAllUsers() ([]domain.User, error) {
	var list []domain.User
	var user domain.User

	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.Password); err != nil {
			return nil, err
		}
		list = append(list, user)
	}
	rows.Close()
	return list, nil
}

func (s *store) CreateUser(user domain.User) error {

	st, err := s.db.Prepare("INSERT INTO users (username, role, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer st.Close()

	res, err := st.Exec(user.Username, user.Role, user.Password)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()

	if err != nil {
		return err
	}

	return nil

}
func (s *store) ExistsUserByUsername(username string) bool {
	var name string
	row := s.db.QueryRow("SELECT username FROM users WHERE username=?", username)

	if err := row.Scan(&name); err != nil {
		return false
	}

	if name == username {
		return true
	}

	return false
}
func (s *store) GetUser(username string) (dto.UserGet, error) {
	var userget dto.UserGet
	row := s.db.QueryRow("SELECT username, role FROM users WHERE username=?", username)
	if err := row.Scan(&userget.Username, &userget.Role); err != nil {
		return dto.UserGet{}, err
	}
	return userget, nil
}

func (s *store) ReadEventById(id int) (domain.Event, error) {
	//defer func () {s.db.Close()}()
	var event domain.Event
	row := s.db.QueryRow("SELECT * FROM eventable WHERE id=?", id)

	if err := row.Scan(&event.Id, &event.Titulo, &event.ShortDescription, &event.LongDescription, &event.State, &event.FechaYHora); err != nil {
		return domain.Event{}, err
		//panic(patient
	}
	return event, nil
}

func (s *store) ReadAllEvents() ([]domain.Event, error) {
	var list []domain.Event

	rows, err := s.db.Query("SELECT * FROM eventable")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event domain.Event
		if err := rows.Scan(&event.Id, &event.Titulo, &event.ShortDescription, &event.LongDescription, &event.State, &event.FechaYHora); err != nil {
			return nil, err
		}
		list = append(list, event)
	}
	rows.Close()
	return list, nil
}

func (s *store) ReadAllEventsWithState() ([]domain.Event, error) {
	var list []domain.Event
	var event domain.Event
	rows, err := s.db.Query("SELECT * FROM eventable WHERE state=?", "PUBLISHED")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&event.Id, &event.Titulo, &event.ShortDescription, &event.LongDescription, &event.State, &event.FechaYHora); err != nil {
			return nil, err
		}
		list = append(list, event)
	}
	rows.Close()
	return list, nil
}

func (s *store) CreateEvent(event domain.Event) error {

	st, err := s.db.Prepare("INSERT INTO eventable (titulo, shortdescription, longdescription, state, fechayhora) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer st.Close()

	res, err := st.Exec(event.Titulo, event.ShortDescription, event.LongDescription, event.State, event.FechaYHora)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
func (s *store) DeleteEvent(id int) error {
	//Preguntar si esta bien usar un metodo
	var idselect int
	row := s.db.QueryRow("SELECT id FROM eventable WHERE id = ?", id)
	if err := row.Scan(&idselect); err != nil {
		return errors.New("The event doest exists.")
	}
	query := "DELETE FROM eventable WHERE id = ?"
	_, err := s.db.Exec(query, idselect)
	if err != nil {
		return err
	}
	return nil

}
func (s *store) UpdateEvent(id int, event domain.Event) error {
	st, err := s.db.Prepare("UPDATE eventable SET titulo = ?, shortdescription = ?, longdescription = ?, state = ?, fechayhora = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer st.Close()

	res, err := st.Exec(event.Titulo, event.ShortDescription, event.LongDescription, event.State, event.FechaYHora, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
