package userHandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"pet/mw"

	"pet/repository/userRepo"
	"pet/service/userService"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func RegisterUserHandlers(r *mux.Router, service *userService.UserService) {
	handler := NewUserHandler(service)

	r.Handle("/users", mw.TokenAuthMiddleware(http.HandlerFunc(handler.GetHandler))).Methods("GET")
	r.Handle("/users", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PutHandler))).Methods("PUT")
	r.Handle("/users", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PostNewUserHandler))).Methods("POST")
	r.Handle("/login", mw.TokenAuthMiddleware(http.HandlerFunc(handler.GetLoginHandler))).Methods("GET")
	r.Handle("/logout", mw.TokenAuthMiddleware(http.HandlerFunc(handler.GetLogoutHandler))).Methods("GET")
	r.Handle("/users", mw.TokenAuthMiddleware(http.HandlerFunc(handler.DeleteUserHandler))).Methods("DELETE")
	r.Handle("/users", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PostNewArrayOfUsersHandler))).Methods("POST")
	r.Handle("/users", mw.TokenAuthMiddleware(http.HandlerFunc(handler.PostNewListOfUserHandler))).Methods("PUT")

}

func (h *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	user, err := h.service.Get(name)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	var user userRepo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.Put(user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) PostNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser userRepo.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.PostNewUser(newUser)
	if err != nil {
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetLogin(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetLogoutHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.GetLogout(credentials.Username)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	err := h.service.DeleteUser(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PostNewArrayOfUsersHandler обрабатывает запрос на создание среза новых пользователей
func (h *UserHandler) PostNewArrayOfUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []userRepo.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.PostNewArrayOfUsers(users)
	if err != nil {
		http.Error(w, "Failed to create new users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) PostNewListOfUserHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON-запрос в список User
	var users []userRepo.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	err = h.service.PostNewListOfUser(users...)
	if err != nil {
		http.Error(w, "Failed to create new users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
