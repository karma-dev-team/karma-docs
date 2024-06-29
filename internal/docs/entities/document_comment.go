package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
)

type DocumentComment struct {
	gormplugin.Model
	DocumentID uuid.UUID `gorm:"type:uuid"`
	Text       string
	MadeBy     uuid.UUID `gorm:"type:uuid"`
	CreatedAt  time.Time
}

func NewDocumentComment(text string, documentId uuid.UUID, madeby uuid.UUID) *DocumentComment {
	return &DocumentComment{
		Text:       text,
		DocumentID: documentId,
		MadeBy:     madeby,
		CreatedAt:  time.Now(),
	}
}
