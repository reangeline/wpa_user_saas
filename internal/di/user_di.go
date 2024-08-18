package di

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/reangeline/wpa_user_saas/internal/domain/usecase"
	"github.com/reangeline/wpa_user_saas/internal/presentation/handler"

	database "github.com/reangeline/wpa_user_saas/internal/infra/database/repository"
)

func InitializeUser(vc *dynamodb.Client) (*handler.UserHanlder, error) {
	userRepository := database.NewUserRepository(vc)

	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
	getUserByPhoneUseCase := usecase.NewGetUserByPhoneUseCase(userRepository)

	user := handler.NewUserHandler(createUserUseCase, getUserByPhoneUseCase)

	return user, nil
}

func InitializeGetUser(vc *dynamodb.Client) (*usecase.GetUserByPhoneUseCase, error) {
	userRepository := database.NewUserRepository(vc)

	getUserByPhoneUseCase := usecase.NewGetUserByPhoneUseCase(userRepository)

	return getUserByPhoneUseCase, nil
}
