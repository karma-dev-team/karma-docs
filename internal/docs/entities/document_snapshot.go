package entities

import (
	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
)

type DocumentSnapshot struct {
	gormplugin.Model
	DocumentId uuid.UUID
	Text       string
	MadeBy     uuid.UUID
}
