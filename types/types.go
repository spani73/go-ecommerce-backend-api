package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User,error)
	GetUserByID(id int ) (*User ,error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"` // this is json marshalling.
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}