package usecase

import (
	"context"
	"errors"

	repository "github.com/reangeline/wpa_user_saas/internal/domain/contract/repository"
	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
	"github.com/reangeline/wpa_user_saas/internal/dto"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewCreateUserUseCase(
	userRepository repository.UserRepositoryInterface,

) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exist")
)

func (u *CreateUserUseCase) Execute(ctx context.Context, input *dto.UserInput) error {

	// userExist, err := u.userRepository.GetUserByEmail(ctx, input.Email)
	// if err != nil {
	// 	return err
	// }

	// if userExist != nil {
	// 	return ErrEmailAlreadyExists
	// }

	user, err := entity.NewUser(input.Name, input.Email, input.PhoneNumber)
	if err != nil {
		return err
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
