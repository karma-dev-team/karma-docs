package entities

import (
	"os/user"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
)

type DocumentSnapshot struct {
	gormplugin.Model
	DocumentId uuid.UUID `gorm:"type:uuid"`
	Text       string
	MadeById   uuid.UUID `gorm:"type:uuid"`
	MadeByUser user.User `gorm:"foreignKey:MadeById;references:ID"`
}

func NewDocumentSnapshot(documentId uuid.UUID, text string, madeby uuid.UUID) *DocumentSnapshot {
	return &DocumentSnapshot{
		DocumentId: documentId,
		Text:       text,
		MadeById:   madeby,
	}
}
