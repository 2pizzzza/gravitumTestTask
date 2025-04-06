package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
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

	err := r.CheckUserExist(ctx, params.Username, params.Email)
	if err != nil {
		return err
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
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

	err := r.CheckUserExist(ctx, params.Username, params.Email)
	if err != nil {
		return err
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
	rowsAffected, err := r.queries.DeleteUser(ctx, int32(id))
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return user.ErrUserNotFound
	}
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	dbUser, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
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

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	dbUser, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
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

func (r *UserRepository) CheckUserExist(ctx context.Context, username, email string) error {
	existingUser, err := r.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return user.ErrUserAlreadyExists
	}

	existingUser, err = r.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return user.ErrUserAlreadyExists
	}
	return nil
}
