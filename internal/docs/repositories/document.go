package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/karma-dev-team/karma-docs/internal/docs/entities"
)

type DocumentRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewDocumentRepository(pool *pgxpool.Pool) *DocumentRepositoryImpl {
	return &DocumentRepositoryImpl{
		pool: pool,
	}
}

func (dr *DocumentRepositoryImpl) AddDocument(ctx context.Context, document *entities.Document) error {
	query := ` 
		INSERT INTO documents(id, )
	`
}
