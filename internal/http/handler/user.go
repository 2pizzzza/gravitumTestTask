package handler

import (
	"errors"
	"net/http"
	"strconv"
	"testTaskGravitum/internal/domain/user"
	"testTaskGravitum/internal/utils"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (uh *UserHandler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /user", uh.createUser)
	mux.HandleFunc("GET /user", uh.getUser)
	mux.HandleFunc("PUT /user", uh.updateUser)
	mux.HandleFunc("DELETE /user", uh.deleteUser)

	return mux
}

// @Summary Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body user.CreateDTO true "User data"
// @Success 200 {object} user.User
// @Failure 400 {object} ErrorMessage
// @Router /user [post]
func (uh *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var dto user.CreateDTO
	if err := utils.ReadRequestBody(r, &dto); err != nil {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid request body"}, http.StatusBadRequest)
		return
	}

	data, err := uh.userService.CreateUser(r.Context(), &dto)
	if err != nil {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Something is wrong"}, http.StatusBadRequest)
		return
	}

	utils.WriteResponseBody(w, data, http.StatusOK)
}

// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} user.User
// @Failure 400 {object} ErrorMessage
// @Failure 404 {object} ErrorMessage
// @Router /user [get]
func (uh *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid id"}, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid id"}, http.StatusBadRequest)
	}
	data, err := uh.userService.GetUser(r.Context(), int64(id))
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			utils.WriteResponseBody(w, ErrorMessage{Message: err.Error()}, http.StatusNotFound)
			return
		}
		utils.WriteResponseBody(w, ErrorMessage{Message: "Something is wrong"}, http.StatusBadRequest)
		return
	}
	utils.WriteResponseBody(w, data, http.StatusOK)
}

// @Summary Update user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Param user body user.UpdateDTO true "Updated user data"
// @Success 200 {object} user.User
// @Failure 400 {object} ErrorMessage
// @Failure 404 {object} ErrorMessage
// @Router /user [put]
func (uh *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid id"}, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid id"}, http.StatusBadRequest)
	}
	var dto user.UpdateDTO
	if err := utils.ReadRequestBody(r, &dto); err != nil {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid request body"}, http.StatusBadRequest)
		return
	}

	data, err := uh.userService.UpdateUser(r.Context(), int64(id), &dto)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			utils.WriteResponseBody(w, ErrorMessage{Message: err.Error()}, http.StatusNotFound)
			return
		}
		utils.WriteResponseBody(w, ErrorMessage{Message: "Something is wrong"}, http.StatusBadRequest)
		return
	}
	utils.WriteResponseBody(w, data, http.StatusOK)
}

// @Summary Delete user by ID
// @Tags Users
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorMessage
// @Failure 404 {object} ErrorMessage
// @Router /user [delete]
func (uh *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid id"}, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid id"}, http.StatusBadRequest)
		return
	}
	err = uh.userService.DeleteUser(r.Context(), int64(id))
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			utils.WriteResponseBody(w, ErrorMessage{Message: err.Error()}, http.StatusNotFound)
			return
		}
		utils.WriteResponseBody(w, ErrorMessage{Message: "Something is wrong"}, http.StatusBadRequest)
		return
	}
	utils.WriteResponseBody(w, struct {
		Message string `json:"message"`
	}{
		Message: "Successes",
	}, http.StatusOK)
}

func (uh *UserHandler) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.WriteResponseBody(w, ErrorMessage{Message: "Invalid email"}, http.StatusBadRequest)
		return
	}

	data, err := uh.userService.GetByEmail(r.Context(), email)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			utils.WriteResponseBody(w, ErrorMessage{Message: err.Error()}, http.StatusNotFound)
			return
		}
		utils.WriteResponseBody(w, ErrorMessage{Message: "Something is wrong"}, http.StatusBadRequest)
		return
	}
	utils.WriteResponseBody(w, data, http.StatusOK)

}
