package usecases

import (
	"github.com/karma-dev-team/karma-docs/internal/docs"
	"github.com/karma-dev-team/karma-docs/internal/docs/repositories"
)

type GroupServiceImpl struct {
	repo repositories.DocumentRepository
}

func NewGroupService(repo repositories.DocumentRepository) docs.GroupService {
	return &GroupServiceImpl{repo: repo}
}
