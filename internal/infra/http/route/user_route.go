package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/wpa_user_saas/internal/presentation/handler"
)

func InitializeUserRoutes(handler *handler.UserHanlder, r chi.Router) {

	r.Route("/user", func(r chi.Router) {
		r.Post("/", handler.CreateUser)
		r.Get("/{phone}", handler.GetUserByPhone)
	})

}
