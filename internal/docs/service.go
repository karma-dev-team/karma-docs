package docs

import (
	"context"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
)

type DocumentService interface {
	CreateDocument(ctx context.Context, dto CreateDocumentDto) (uuid.UUID, error)
	GetDocument(ctx context.Context, documentId uuid.UUID) (*entities.Document, error)
	UpdateDocument(ctx context.Context, dto UpdateDocumentDto) error
	DeleteDocument(ctx context.Context, documentId uuid.UUID) error
	GetDocumentsList(ctx context.Context, dto GetDocumentsListDto) ([]*entities.Document, error)
}

type GroupService interface {
	CreateGroup(ctx context.Context, dto CreateGroupDto) (uuid.UUID, error)
	GetGroup(ctx context.Context, groupId uuid.UUID) (*entities.Group, error)
	GetGroupsList(ctx context.Context, userId uuid.UUID) ([]*entities.Group, error)
	InviteMember(ctx context.Context, dto InviteUserDto) error
	DeleteGroup(ctx context.Context, groupId uuid.UUID) error
}
