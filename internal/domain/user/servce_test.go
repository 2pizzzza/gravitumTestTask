package user_test

import (
	"context"
	"errors"
	userService "testTaskGravitum/internal/service/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testTaskGravitum/internal/domain/user"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(ctx context.Context, u *user.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepo) GetByID(ctx context.Context, id int64) (*user.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockRepo) Update(ctx context.Context, u *user.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepo) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

// ====== Тесты ======

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	dto := &user.CreateDTO{
		Username: "john",
		Email:    "john@example.com",
	}

	expected := &user.User{
		Username: dto.Username,
		Email:    dto.Email,
	}

	mockRepo.On("Create", mock.Anything, expected).Return(nil)

	result, err := svc.CreateUser(context.Background(), dto)

	assert.NoError(t, err)
	assert.Equal(t, expected.Username, result.Username)
	assert.Equal(t, expected.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_ValidationError(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	dto := &user.CreateDTO{}

	result, err := svc.CreateUser(context.Background(), dto)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	id := int64(1)
	existing := &user.User{
		ID:       id,
		Username: "old",
		Email:    "old@example.com",
	}
	updated := &user.User{
		ID:       id,
		Username: "new",
		Email:    "new@example.com",
	}

	dto := &user.UpdateDTO{
		Username: "new",
		Email:    "new@example.com",
	}

	mockRepo.On("GetByID", mock.Anything, id).Return(existing, nil)
	mockRepo.On("Update", mock.Anything, updated).Return(nil)

	result, err := svc.UpdateUser(context.Background(), id, dto)

	assert.NoError(t, err)
	assert.Equal(t, dto.Username, result.Username)
	assert.Equal(t, dto.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	id := int64(42)
	dto := &user.UpdateDTO{Username: "x"}

	mockRepo.On("GetByID", mock.Anything, id).Return(nil, errors.New("not found"))

	result, err := svc.UpdateUser(context.Background(), id, dto)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	id := int64(10)
	mockRepo.On("Delete", mock.Anything, id).Return(nil)

	err := svc.DeleteUser(context.Background(), id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Error(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	id := int64(10)
	mockRepo.On("Delete", mock.Anything, id).Return(errors.New("delete error"))

	err := svc.DeleteUser(context.Background(), id)

	assert.Error(t, err)
}

func TestGetUser_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	id := int64(5)
	expected := &user.User{ID: id, Username: "test", Email: "test@example.com"}

	mockRepo.On("GetByID", mock.Anything, id).Return(expected, nil)

	result, err := svc.GetUser(context.Background(), id)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	id := int64(5)
	mockRepo.On("GetByID", mock.Anything, id).Return(nil, errors.New("not found"))

	result, err := svc.GetUser(context.Background(), id)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetByEmail_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	email := "test@example.com"
	expected := &user.User{ID: 1, Username: "test", Email: email}

	mockRepo.On("GetByEmail", mock.Anything, email).Return(expected, nil)

	result, err := svc.GetByEmail(context.Background(), email)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetByEmail_NotFound(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := userService.New(mockRepo)

	email := "missing@example.com"
	mockRepo.On("GetByEmail", mock.Anything, email).Return(nil, errors.New("not found"))

	result, err := svc.GetByEmail(context.Background(), email)

	assert.Error(t, err)
	assert.Nil(t, result)
}
