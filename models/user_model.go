package models

type User struct {
	Username 	string			   `json:"username,omitempty" validate:"required"`
	Password	string			   `json:"password,omitempty" validate:"required"`
}	