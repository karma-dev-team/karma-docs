package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
)

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: pool}
}

func (repo *UserRepositoryImpl) AddUser(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users (id, username, email)
		VALUES ($1, $2, $3)
	`
	_, err := repo.db.Exec(ctx, query, user.Id, user.Username, user.Email)
	return err
}

func (repo *UserRepositoryImpl) GetUser(ctx context.Context, request GetUserRequest) (*entities.User, error) {
	query := `
		SELECT id, username, email
		FROM users
		WHERE id = $1 OR username = $2 OR email = $3
	`
	row := repo.db.QueryRow(ctx, query, request.UserId, request.Username, request.Email)

	user := &entities.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err == pgx.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	_, err := repo.db.Exec(ctx, query, userId)
	return err
}
