package user

import (
	"context"
	"errors"
	"testTaskGravitum/internal/domain/user"
)

func (s *Service) CreateUser(ctx context.Context, dto *user.CreateDTO) (*user.User, error) {
	if dto.Username == "" || dto.Email == "" {
		return nil, errors.New("username and email are required")
	}
	userRaw := &user.User{
		Username: dto.Username,
		Email:    dto.Email,
	}
	err := s.repo.Create(ctx, userRaw)
	if err != nil {
		return nil, err
	}
	return userRaw, nil
}

func (s *Service) UpdateUser(ctx context.Context, id int64, dto *user.UpdateDTO) (*user.User, error) {
	userRaw, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if dto.Username != "" {
		userRaw.Username = dto.Username
	}
	if dto.Email != "" {
		userRaw.Email = dto.Email
	}
	err = s.repo.Update(ctx, userRaw)
	if err != nil {
		return nil, err
	}
	return userRaw, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int64) error {

	return nil
}

func (s *Service) GetUser(ctx context.Context, id int64) (*user.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*user.User, error) {
	return nil, nil
}
