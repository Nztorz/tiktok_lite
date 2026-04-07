package services

import (
	"context"
	"fmt"

	"github.com/Nztorz/tiktok_lite/internal/auth"
	"github.com/Nztorz/tiktok_lite/internal/database"
	"github.com/Nztorz/tiktok_lite/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

// hash the password and send data to repository to the final regitration in db
func (u *UserService) Register(ctx context.Context, email, username, password, bio string) (database.User, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return database.User{}, fmt.Errorf("error hashing %w", err)
	}

	// register to db
	user, err := u.repo.Create(ctx, email, username, hash, bio)
	if err != nil {
		return user, err
	}

	return user, nil
}
