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
	ReadAllUsers() ([]domain.User, error)
	CreateUser(user domain.User) error
	ExistsUserByUsername(username string) bool
	//------ EVENT

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
	row := s.db.QueryRow("SELECT username FROM users WHERE username=?",username)

	if err := row.Scan(&name); err != nil {
		return false
	}

	if name == username {
		return true
	}

	return false
}
