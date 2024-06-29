package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
	"gorm.io/gorm"
)

type DocumentRepositoryImpl struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepositoryImpl {
	return &DocumentRepositoryImpl{
		db: db,
	}
}

func (r *DocumentRepositoryImpl) AddDocument(ctx context.Context, document *entities.Document) error {
	result := r.db.WithContext(ctx).Create(document)
	return result.Error
}

func (r *DocumentRepositoryImpl) EditDocument(ctx context.Context, document *entities.Document) error {
	result := r.db.WithContext(ctx).Save(document)
	return result.Error
}

func (r *DocumentRepositoryImpl) GetDocument(ctx context.Context, documentId uuid.UUID) (*entities.Document, error) {
	var document entities.Document
	result := r.db.WithContext(ctx).Preload("Snapshots").Preload("Comments").First(&document, "id = ?", documentId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &document, result.Error
}

func (r *DocumentRepositoryImpl) DeleteDocument(ctx context.Context, documentId uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&entities.Document{}, "id = ?", documentId)
	return result.Error
}
