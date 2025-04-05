package user

import (
	"testTaskGravitum/internal/domain/user"
)

type Service struct {
	repo user.Repository
}

func New(repo user.Repository) *Service {
	return &Service{repo: repo}
}
