package di

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/reangeline/wpa_user_saas/internal/domain/usecase"
	database "github.com/reangeline/wpa_user_saas/internal/infra/database/repository"
	sns_impl "github.com/reangeline/wpa_user_saas/internal/infra/event/implementation"
	"github.com/reangeline/wpa_user_saas/internal/presentation/controller"
	"github.com/reangeline/wpa_user_saas/internal/presentation/handler"
)

func InitializeUser(vc *dynamodb.DynamoDB) (*controller.UserController, error) {
	userRepository := database.NewUserRepository(vc)

	snsService := sns_impl.NewSNSService("arn:aws:sns:us-east-1:237071355172:GoWpaUser")

	createUserUseCase := usecase.NewCreateUserUseCase(userRepository, snsService)
	userController := controller.NewUserController(createUserUseCase)
	return userController, nil
}

func InitializeEventUser(vc *dynamodb.DynamoDB) (*handler.UserHandler, error) {
	userRepository := database.NewUserRepository(vc)

	snsService := sns_impl.NewSNSService("arn:aws:sns:us-east-1:237071355172:GoWpaUser")

	createUserUseCase := usecase.NewCreateUserUseCase(userRepository, snsService)

	userHandler := handler.NewUserHandler(createUserUseCase)

	return userHandler, nil
}
