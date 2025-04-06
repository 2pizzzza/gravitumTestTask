package handler

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "testTaskGravitum/docs"
	service "testTaskGravitum/internal/service/user"
)

type Handler struct {
	User *UserHandler
}

func New(services *service.Service) *Handler {
	return &Handler{
		User: NewUserHandler(services),
	}
}

func (h *Handler) RegisterRouter(mux *http.ServeMux) {
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.Handle("/", h.User.Router())

}
