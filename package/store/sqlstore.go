package store

import ("database/sql")

type store struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) Store {
	return &store{db: db}
}

type Store interface {

	
}