package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/auth"
	"github.com/karma-dev-team/karma-docs/internal/docs"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
	"github.com/karma-dev-team/karma-docs/internal/docs/repositories"
	. "github.com/openfga/go-sdk/client"
)

type DocumentServiceImpl struct {
	repo      repositories.DocumentRepository
	fgaClient SdkClient // it's impossible to implement general interface to not to
}

func NewDocumentService(repo repositories.DocumentRepository, fgaClient SdkClient) *DocumentServiceImpl {
	return &DocumentServiceImpl{
		repo:      repo,
		fgaClient: fgaClient,
	}
}

func (s *DocumentServiceImpl) CreateDocument(ctx context.Context, dto docs.CreateDocumentDto) (uuid.UUID, error) {
	if dto.GroupId != nil {
		body := ClientCheckRequest{
			User:     "user:" + dto.AuthorId.Version().String(),
			Relation: "can_add_document",
			Object:   "group:" + dto.GroupId.Version().String(),
		}

		resp, err := s.fgaClient.Check(ctx).Body(body).Execute()
		if err != nil {
			return uuid.Nil, err
		}
		if !*resp.Allowed {
			return uuid.Nil, auth.ErrAccessDenied
		}
	}
	document := entities.NewDocument(dto.Title, dto.AuthorId, dto.Text)
	err := s.repo.AddDocument(ctx, document)
	if err != nil {
		return uuid.Nil, err
	}

	return document.ID, nil
}

func (s *DocumentServiceImpl) GetDocument(ctx context.Context, documentId uuid.UUID) (*entities.Document, error) {
	body := ClientCheckRequest{}
	return s.repo.GetDocument(ctx, documentId)
}

func (s *DocumentServiceImpl) UpdateDocument(ctx context.Context, dto docs.UpdateDocumentDto) error {
	document := &entities.Document{
		// Initialize fields from dto
	}
	return s.repo.EditDocument(ctx, document)
}

func (s *DocumentServiceImpl) DeleteDocument(ctx context.Context, documentId uuid.UUID) error {
	return s.repo.DeleteDocument(ctx, documentId)
}

func (s *DocumentServiceImpl) GetDocumentsList(ctx context.Context, dto docs.GetDocumentsListDto) ([]*entities.Document, error) {
	// Implement logic to fetch a list of documents based on dto criteria
	// Example:
	documents, err := s.repo.GetDocumentsList(ctx, repositories.GetDocumentsListQuery{
		AuthorId: dto.AuthorId,
		GroupId:  dto.GroupId,
	})
	if err != nil {
		return nil, err
	}
	return documents, nil
}
