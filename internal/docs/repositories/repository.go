package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
)

type DocumentRepository interface {
	AddDocument(ctx context.Context, document *entities.Document) error
	EditDocument(ctx context.Context, document *entities.Document) error
	GetDocument(ctx context.Context, documentId uuid.UUID) (*entities.Document, error)
	DeleteDocument(ctx context.Context, documentId uuid.UUID) error
}
