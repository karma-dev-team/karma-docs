package usecases

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/auth"
	"github.com/karma-dev-team/karma-docs/internal/config"
	"github.com/karma-dev-team/karma-docs/internal/docs"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
	"github.com/karma-dev-team/karma-docs/internal/docs/repositories"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
	. "github.com/openfga/go-sdk/client"
)

type DocumentServiceImpl struct {
	repo      repositories.DocumentRepository
	fgaClient SdkClient // it's impossible to implement general interface to not to
	config    *config.AppConfig
}

func NewDocumentService(repo repositories.DocumentRepository, fgaClient SdkClient, config *config.AppConfig) *DocumentServiceImpl {
	return &DocumentServiceImpl{
		repo:      repo,
		fgaClient: fgaClient,
		config:    config,
	}
}

func (s *DocumentServiceImpl) CreateDocument(ctx context.Context, dto docs.CreateDocumentDto) (uuid.UUID, error) {
	if dto.GroupId != nil {
		req := ClientCheckRequest{
			User:     "user:" + dto.AuthorId.Version().String(),
			Relation: "can_add_document",
			Object:   "group:" + dto.GroupId.Version().String(),
		}

		resp, err := s.fgaClient.
			Check(ctx).
			Body(req).
			Options(ClientCheckOptions{AuthorizationModelId: &s.config.Openfga.AuthorizationModelId}).
			Execute()
		if err != nil {
			return uuid.Nil, err
		}
		if !*resp.Allowed {
			return uuid.Nil, auth.ErrAccessDenied
		}
	}
	document := entities.NewDocument(dto.Title, dto.AuthorId, dto.Text)
	documentId, err := s.repo.AddDocument(ctx, document)
	if err != nil {
		return uuid.Nil, err
	}

	addReq := ClientWriteRequest{
		Writes: []ClientTupleKey{
			{
				User:     "group:" + dto.GroupId.String(),
				Relation: "write",
				Object:   "document:" + documentId.String(),
			},
		},
	}

	_, err = s.fgaClient.
		Write(ctx).
		Body(addReq).
		Options(ClientWriteOptions{AuthorizationModelId: &s.config.Openfga.AuthorizationModelId}).
		Execute()
	if err != nil {
		return uuid.Nil, err
	}

	return document.ID, nil
}

func (s *DocumentServiceImpl) GetDocument(ctx context.Context, documentId uuid.UUID, byUser uuid.UUID) (*entities.Document, error) {
	body := ClientCheckRequest{
		User:     "group:" + byUser.String(),
		Object:   "document" + documentId.String(),
		Relation: "read",
	}
	resp, err := s.fgaClient.
		Check(ctx).
		Body(body).
		Options(ClientCheckOptions{AuthorizationModelId: &s.config.Openfga.AuthorizationModelId}).
		Execute()
	if err != nil {
		return nil, err
	}
	if !*resp.Allowed {
		return nil, auth.ErrAccessDenied
	}
	return s.repo.GetDocument(ctx, documentId)
}

func (s *DocumentServiceImpl) UpdateDocument(ctx context.Context, dto docs.UpdateDocumentDto) error {
	err := s.canWriteToDocument(ctx, dto.ByGroup, dto.DocumentID)
	if err != nil {
		return err
	}

	document := &entities.Document{
		Model: gormplugin.Model{ID: dto.DocumentID},
		Title: dto.Title,
		Text:  dto.Text,
	}

	return s.repo.EditDocument(ctx, document)
}

func (s *DocumentServiceImpl) canWriteToDocument(ctx context.Context, byGroupId uuid.UUID, documentId uuid.UUID) error {
	body := ClientCheckRequest{
		User:     "group:" + byGroupId.String(),
		Relation: "write",
		Object:   "document:" + documentId.String(),
	}

	resp, err := s.fgaClient.
		Check(ctx).
		Body(body).
		Options(ClientCheckOptions{AuthorizationModelId: &s.config.Openfga.AuthorizationModelId}).
		Execute()

	if err != nil {
		return err
	}

	if !*resp.Allowed {
		return auth.ErrAccessDenied
	}
	return nil
}

func (s *DocumentServiceImpl) DeleteDocument(ctx context.Context, documentId uuid.UUID, groupId uuid.UUID) error {
	err := s.canWriteToDocument(ctx, groupId, documentId)
	if err != nil {
		return err
	}

	return s.repo.DeleteDocument(ctx, documentId)
}

func (s *DocumentServiceImpl) GetDocumentsList(ctx context.Context, dto docs.GetDocumentsListDto) ([]*entities.Document, error) {
	resp, err := s.fgaClient.ListObjects(ctx).Body(ClientListObjectsRequest{
		User:     "group:" + dto.GroupId.String(),
		Relation: "read",
		Type:     "document",
	}).Options(
		ClientListObjectsOptions{AuthorizationModelId: &s.config.Openfga.AuthorizationModelId},
	).Execute()
	if err != nil {
		return nil, err
	}
	var documentIds []uuid.UUID
	for _, obj := range resp.GetObjects() {
		id := strings.Split(obj, ":")[1]
		documentId, err := uuid.FromBytes([]byte(id))
		if err != nil {
			return nil, err
		}
		documentIds = append(documentIds, documentId)
	}

	documents, err := s.repo.GetDocumentsList(ctx, repositories.GetDocumentsListQuery{
		GroupId:     dto.GroupId,
		DocumentIds: documentIds,
	})
	if err != nil {
		return nil, err
	}
	return documents, nil
}
