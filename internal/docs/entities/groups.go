package entities

import (
	"github.com/google/uuid"
	"github.com/karma-dev-team/karma-docs/internal/user"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
	"github.com/karma-dev-team/karma-docs/pkg/gormplugin"
)

// Does not have any owned Documents, because ownership defined in
// permission model, not in database, thus way making it much simpler
//
// And yes, i know that will cause problem when changing permission model and etc.
// and it has another problem, that database becames anemic lol, idc
type Group struct {
	gormplugin.Model
	Name        string
	Users       []*entities.User `gorm:"many2many:user_groups"`
	Description string
	CreatedBy   entities.User `gorm:"foreignkey:ID"`
}

func NewGroups(name, description string, users []*entities.User) *Group {
	return &Group{
		Name:        name,
		Description: description,
		Users:       users,
	}
}

func (g *Group) AddUser(newUser *entities.User) error {
	if newUser.IsBlocked {
		return user.ErrUserIsBlocked
	}
	for _, groupUser := range g.Users {
		if groupUser.ID == newUser.ID {
			return nil
		}
	}
	g.Users = append(g.Users, newUser)
	return nil
}

func (g *Group) RemoveUser(userID uuid.UUID) {
	for i, groupUser := range g.Users {
		if groupUser.ID == userID {
			g.Users = append(g.Users[:i], g.Users[i+1:]...)
			break
		}
	}
}
