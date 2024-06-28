package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
	"github.com/karma-dev-team/karma-docs/internal/user/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, dto user.CreateUserDto) error {
	// Create a new user entity
	user, err := entities.NewUser(dto.Username, dto.Email, dto.Password)
	if err != nil {
		return err
	}

	// Add the user to the repository
	return s.repo.AddUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, userId uuid.UUID) (*entities.User, error) {
	request := repositories.GetUserRequest{
		UserId: userId,
	}

	// Get the user from the repository
	user, err := s.repo.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	// Delete the user from the repository
	return s.repo.DeleteUser(ctx, userId)
}
