package domain

import ()

type Inscription struct {
	Id    int   `json:"id"`
	Event Event `json:"event,omitempty" xml:"event,omitempty"`
	User  User  `json:"user,omitempty" xml:"user,omitempty"`
}
