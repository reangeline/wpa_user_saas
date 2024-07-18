package handler

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/reangeline/wpa_user_saas/internal/domain/contract/usecase"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	createUserUseCase usecase.CreateUserUseCaseInterface
}

func NewUserHandler(
	createUserUseCase usecase.CreateUserUseCaseInterface,
) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (uh *UserHandler) CreateUserLambda(e events.SNSEntity) {

	fmt.Println(e.Message, e.Subject, "aqui")
	// Do something
}
