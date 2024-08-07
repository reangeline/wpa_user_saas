package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/reangeline/wpa_user_saas/internal/domain/contract/repository/mock"
	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
	"github.com/reangeline/wpa_user_saas/internal/dto"
	pkg "github.com/reangeline/wpa_user_saas/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepositoryInterface(ctrl)
	useCase := NewCreateUserUseCase(mockRepo)

	t.Run("should create user successfully", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, nil)

		input := &dto.UserInput{
			Name:     "John",
			LastName: "Doe",
			Email:    "john.doe@example.com",
			Phone:    "1234567890",
		}

		err := useCase.Execute(context.Background(), input)

		assert.NoError(t, err)
	})

	t.Run("should return error when email already exists", func(t *testing.T) {
		input := &dto.UserInput{
			Name:     "Jane",
			LastName: "Doe",
			Email:    "jane.doe@example.com",
			Phone:    "0987654321",
		}

		mockRepo.EXPECT().GetUserByEmail(gomock.Any(), input.Email).Return(
			&entity.User{
				IDUser:   pkg.NewID(),
				Name:     input.Name,
				LastName: input.LastName,
				Email:    input.Email,
				Phone:    input.Phone,
			}, nil)

		err := useCase.Execute(context.Background(), input)

		assert.Error(t, err)
		assert.Equal(t, ErrEmailAlreadyExists, err)
	})

	t.Run("should return error when user creation fails", func(t *testing.T) {
		input := &dto.UserInput{
			Name:     "Mark",
			LastName: "Smith",
			Email:    "mark.smith@example.com",
			Phone:    "1112223333",
		}

		mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("creation failed"))

		mockRepo.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(nil, nil)

		err := useCase.Execute(context.Background(), input)

		assert.Error(t, err)
		assert.Equal(t, "creation failed", err.Error())
	})
}
