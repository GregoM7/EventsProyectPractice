package domain

type User struct {
	Id int `json:"id"`
	Username string `json:"username,omitempty"`
}
