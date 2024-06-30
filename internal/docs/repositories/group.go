package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
	"gorm.io/gorm"
)

type GroupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &GroupRepositoryImpl{db: db}
}

func (repo *GroupRepositoryImpl) AddGroup(ctx context.Context, group *entities.Group) error {
	return repo.db.WithContext(ctx).Create(group).Error
}

// EditGroup updates an existing group in the database
func (repo *GroupRepositoryImpl) EditGroup(ctx context.Context, group *entities.Group) error {
	return repo.db.WithContext(ctx).Save(group).Error
}

// GetGroup retrieves a group from the database by its ID
func (repo *GroupRepositoryImpl) GetGroup(ctx context.Context, groupId uuid.UUID) (*entities.Group, error) {
	var group entities.Group
	if err := repo.db.WithContext(ctx).Preload("Users").First(&group, "id = ?", groupId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &group, nil
}

// DeleteGroup deletes a group from the database by its ID
func (repo *GroupRepositoryImpl) DeleteGroup(ctx context.Context, groupId uuid.UUID) error {
	return repo.db.WithContext(ctx).Delete(&entities.Group{}, "id = ?", groupId).Error
}
