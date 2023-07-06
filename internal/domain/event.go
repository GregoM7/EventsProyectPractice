package domain

import (
	

	"time"
)

type Event struct {
	Id               int       `json:"idevents"`
	Titulo           string    `json:"titulo,omitempty"`
	ShortDescription string    `json:"shortdescription,omitempty"`
	LongDescription  string    `json:"longdescription,omitempty"`
	State            string    `json:"state,omitempty"`
	FechaYHora       time.Time `json:"fechayhora,omitempty"`
}
