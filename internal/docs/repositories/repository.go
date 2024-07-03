package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
)

type GetDocumentsListQuery struct {
	GroupId     uuid.UUID
	DocumentIds []uuid.UUID
}

type DocumentRepository interface {
	AddDocument(ctx context.Context, document *entities.Document) (uuid.UUID, error)
	EditDocument(ctx context.Context, document *entities.Document) error
	GetDocument(ctx context.Context, documentId uuid.UUID) (*entities.Document, error)
	GetDocumentsList(ctx context.Context, query GetDocumentsListQuery) ([]*entities.Document, error)
	DeleteDocument(ctx context.Context, documentId uuid.UUID) error
}

type GroupRepository interface {
	AddGroup(ctx context.Context, group *entities.Group) error
	EditGroup(ctx context.Context, group *entities.Group) error
	GetGroup(ctx context.Context, groupId uuid.UUID) (*entities.Group, error)
	DeleteGroup(ctx context.Context, groupId uuid.UUID) error
}
