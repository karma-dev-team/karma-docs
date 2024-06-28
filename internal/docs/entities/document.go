package entities

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id        uuid.UUID
	Title     string
	OwnerId   uuid.UUID
	Text      string
	CreatedAt time.Time
	Version   string // optional
}

type DocumentService struct{}

func NewDocument(
	title string,
	ownerId uuid.UUID,
	text string,
) {

}
