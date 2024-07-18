package usecase

import (
	"context"

	"github.com/reangeline/wpa_user_saas/internal/dto"
)

type CreateUserUseCaseInterface interface {
	Execute(ctx context.Context, input *dto.UserInput) error
}
