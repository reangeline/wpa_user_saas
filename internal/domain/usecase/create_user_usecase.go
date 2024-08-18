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
	ErrPhoneAlreadyExists = errors.New("phone already exist")
)

func (u *CreateUserUseCase) Execute(ctx context.Context, input *dto.UserInput) error {

	// isExist, err := u.userRepository.GetUserByPhone(ctx, input.PhoneNumber)
	// if err != nil {
	// 	return err
	// }

	// if isExist != nil {
	// 	return ErrPhoneAlreadyExists
	// }

	user, err := entity.NewUser(input.Name, input.Email, input.Phone)
	if err != nil {
		return err
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
