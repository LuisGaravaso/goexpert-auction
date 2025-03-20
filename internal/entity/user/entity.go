package user

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, id string) (*User, *internal_errors.InternalError)
}
