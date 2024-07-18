package contract

import (
	"context"

	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
}

// mockgen -source=internal/domain/usecase/create_user_usecase.go -destination=internal/domain/usecase/mock/create_user_usecase_mock.go -package=usecase
