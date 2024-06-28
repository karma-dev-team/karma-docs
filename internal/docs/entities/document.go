package entities

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id                uuid.UUID
	Title             string
	OwnerId           uuid.UUID
	Text              string
	CreatedAt         time.Time
	CurrentSnapshotId uuid.UUID
	// eagerly loaded
	Snapshots []DocumentSnapshot
}

func NewDocument(
	title string,
	ownerId uuid.UUID,
	text string,
) *Document {
	// ignore error, because we dont use pool randomization, and never will! so it is HIGHLY unlikly to cause any errors
	id, _ := uuid.NewRandom()
	return &Document{
		Id:                id,
		OwnerId:           ownerId,
		Text:              text,
		CurrentSnapshotId: uuid.Nil,
	}
}

func (d *Document) Snapshot(madeById uuid.UUID, newText string) {

}
