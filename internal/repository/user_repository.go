package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Nztorz/tiktok_lite/internal/database"
)

type UserRepository interface {
	Create(ctx context.Context, email, username, hash, bio string) (database.User, error)
}

type userRepository struct {
	q *database.Queries
}

func (u *userRepository) Create(ctx context.Context, email, username, hash, bio string) (database.User, error) {

	user, err := u.q.CreateUser(ctx, database.CreateUserParams{
		Email:          email,
		Username:       username,
		HashedPassword: hash,
		Bio: sql.NullString{
			String: bio,
			Valid:  bio != "",
		},
	})

	if err != nil {
		return database.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
