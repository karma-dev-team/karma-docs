package entities

import "github.com/google/uuid"

type Document struct {
	Id      uuid.UUID
	Name    string
	OwnerId uuid.UUID
}

type DocumentService struct{}
