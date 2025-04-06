package user

import (
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, dto *CreateDTO) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, id int64, dto *UpdateDTO) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
