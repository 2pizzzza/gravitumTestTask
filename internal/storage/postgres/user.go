package postgres

import (
	"context"
	"testTaskGravitum/internal/domain/user"
)

type UserRepository struct {
	queries *Queries
}

func NewUsersRepository(db DBTX) *UserRepository {
	return &UserRepository{
		queries: New(db),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	params := CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
	}
	createdUser, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	user.ID = int64(createdUser.ID)
	user.CreatedAt = createdUser.CreatedAt.Time
	user.UpdatedAt = createdUser.UpdatedAt.Time
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	dbUser, err := r.queries.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:        int64(dbUser.ID),
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:        int64(dbUser.ID),
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}, nil
}

func (r *UserRepository) Update(ctx context.Context, user *user.User) error {
	params := UpdateUserParams{
		ID:       int32(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}
	updatedUser, err := r.queries.UpdateUser(ctx, params)
	if err != nil {
		return err
	}
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.UpdatedAt = updatedUser.UpdatedAt.Time
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.queries.DeleteUser(ctx, int32(id))
}
