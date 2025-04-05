package user

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, dto *CreateDTO) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id int64, dto *UpdateDTO) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
}
