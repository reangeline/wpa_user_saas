package usecase

import (
	"context"
	"errors"

	event "github.com/reangeline/wpa_user_saas/internal/domain/contract/event"
	repository "github.com/reangeline/wpa_user_saas/internal/domain/contract/repository"
	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
	"github.com/reangeline/wpa_user_saas/internal/dto"
)

type CreateUserUseCase struct {
	userRepository repository.UserRepositoryInterface
	snsService     event.SNSServiceInterface
}

func NewCreateUserUseCase(
	userRepository repository.UserRepositoryInterface,
	snsService event.SNSServiceInterface,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
		snsService:     snsService,
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

	user, err := entity.NewUser(input.Name, input.LastName, input.Email, input.Phone)
	if err != nil {
		return err
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	err = u.snsService.PublishMessage("user", "created_user")
	if err != nil {
		return err
	}

	return nil
}
