package usecases

import "github.com/karma-dev-team/karma-docs/internal/docs/repositories"

type DocumentServiceImpl struct {
	repo repositories.DocumentRepository
}

func NewDocumentService(repo repositories.DocumentRepository) *DocumentServiceImpl {
	return &DocumentServiceImpl{
		repo: repo,
	}
}
