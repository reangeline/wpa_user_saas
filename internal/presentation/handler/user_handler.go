package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reangeline/wpa_user_saas/internal/domain/contract/usecase"
	"github.com/reangeline/wpa_user_saas/internal/dto"
)

type UserHanlder struct {
	createUserUseCase     usecase.CreateUserUseCaseInterface
	getUserByPhoneUseCase usecase.GetUserByPhoneUseCaseInterface
}

func NewUserHandler(
	createUserUseCase usecase.CreateUserUseCaseInterface,
	getUserByPhoneUseCase usecase.GetUserByPhoneUseCaseInterface,
) *UserHanlder {
	return &UserHanlder{
		createUserUseCase:     createUserUseCase,
		getUserByPhoneUseCase: getUserByPhoneUseCase,
	}
}

type Error struct {
	Message string `json:"message"`
}

func (uh *UserHanlder) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ctx := r.Context()
	err = uh.createUserUseCase.Execute(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (uh *UserHanlder) GetUserByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.URL.Query().Get("phone")

	ctx := r.Context()
	fmt.Println("phone", phone)
	user, err := uh.getUserByPhoneUseCase.Execute(ctx, phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}
