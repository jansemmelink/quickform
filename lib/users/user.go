package users

import (
	"github.com/satori/go.uuid"
)

//New user
func New() IUser {
	return &user{
		id: uuid.NewV1().String(),
	}
}

//IUser ...
type IUser interface {
	ID() string
	String() string
}

type user struct {
	id string
}

func (u user) ID() string {
	return u.id
}

func (u user) String() string {
	return u.id
}
