package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
)

// aggregate root!!!
type Document struct {
	gormplugin.Model
	Title   string
	OwnerId uuid.UUID
	Text    string
	// eagerly loaded
	Snapshots        []DocumentSnapshot `gorm:"foreignkey:DocumentId"`
	Comments         []DocumentComment  `gorm:"foreignKey:DocumentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LastModifiedDate time.Time
}

func NewDocument(
	title string,
	ownerId uuid.UUID,
	text string,
) *Document {
	return &Document{
		OwnerId:          ownerId,
		Text:             text,
		LastModifiedDate: time.Now(),
		Snapshots:        []DocumentSnapshot{},
	}
}

func (d *Document) ChangeText(text string, madeby uuid.UUID) {
	d.Text = text
	d.Snapshots = append(d.Snapshots, *NewDocumentSnapshot(d.ID, text, madeby))
	d.LastModifiedDate = time.Now()
}

func (d *Document) AddComment(text string, madeby uuid.UUID) {
	d.Comments = append(d.Comments, *NewDocumentComment(text, d.ID, madeby))
}
