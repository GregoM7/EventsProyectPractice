package store

import (
	"database/sql"

	"github.com/GregoM7/EventsProyectPractice/internal/domain"
)

type store struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) Store {
	return &store{db: db}
}

type Store interface {
	//------ USER
	ReadAllUsers()([]domain.User, error)
	//------ EVENT

	//------ INSCRIPTION
}

func (s *store) ReadAllUsers() ([]domain.User,error){
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