package user

import (
	"context"

	"github.com/LuisGaravaso/goexpert-auction/internal/entity/user"
	"github.com/LuisGaravaso/goexpert-auction/internal/internal_errors"
)

type UserInputDTO struct {
	ID string `json:"id"`
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUseCase struct {
	UserRepository user.UserRepositoryInterface
}

type UserUsecaseInterface interface {
	FindUserById(ctx context.Context, input UserInputDTO) (*UserOutputDTO, *internal_errors.InternalError)
}

func NewUserUsecase(userRepository user.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{UserRepository: userRepository}
}

func (u *UserUseCase) FindUserById(ctx context.Context, input UserInputDTO) (*UserOutputDTO, *internal_errors.InternalError) {
	user, err := u.UserRepository.FindUserById(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
