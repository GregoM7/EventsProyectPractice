package domain

import (
	"time"
)

type Event struct {
	Id               int       `json:"id"`
	Titulo           string    `json:"titulo,omitempty"`
	ShortDescription string    `json:"short_description,omitempty"`
	LongDescription  string    `json:"long_description,omitempty"`
	FechaYHora       time.Time `json:"fechayhora,omitempty"`
}
