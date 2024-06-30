package docs

import "github.com/google/uuid"

type CreateDocumentDto struct {
	Title    string
	Text     string
	AuthorId uuid.UUID
	GroupId  *uuid.UUID
}

type UpdateDocumentDto struct {
	Title string
	Text  string
}

type GetDocumentsListDto struct {
	GroupId  uuid.UUID // optional
	AuthorId uuid.UUID // optional
}

type CreateGroupDto struct {
	Name        string
	Description string
}

type InviteUserDto struct {
	GroupId uuid.UUID
	UserId  uuid.UUID
}
