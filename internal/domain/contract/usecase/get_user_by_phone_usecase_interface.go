package usecase

import (
	"context"

	"github.com/reangeline/wpa_user_saas/internal/dto"
)

type GetUserByPhoneUseCaseInterface interface {
	Execute(ctx context.Context, phone string) (*dto.UserOutput, error)
}
