package usecase

import (
	"context"

	repository "github.com/reangeline/wpa_user_saas/internal/domain/contract/repository"
	"github.com/reangeline/wpa_user_saas/internal/dto"
)

type GetUserByPhoneUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewGetUserByPhoneUseCase(
	userRepository repository.UserRepositoryInterface,

) *GetUserByPhoneUseCase {
	return &GetUserByPhoneUseCase{
		userRepository: userRepository,
	}
}

func (u *GetUserByPhoneUseCase) Execute(ctx context.Context, phone string) (*dto.UserOutput, error) {

	user, err := u.userRepository.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	return &dto.UserOutput{
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		UserExist: true,
	}, nil

}
