package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) AddUser(ctx context.Context, user *entities.User) error {
	result := repo.db.WithContext(ctx).Create(user)
	return result.Error
}

func (repo *UserRepositoryImpl) GetUser(ctx context.Context, request GetUserRequest) (*entities.User, error) {
	var user entities.User
	result := repo.db.WithContext(ctx).Where("id = ? OR username = ? OR email = ?", request.UserId, request.Username, request.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}

func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	result := repo.db.WithContext(ctx).Delete(&entities.User{}, userId)
	return result.Error
}
