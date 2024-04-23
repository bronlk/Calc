package main

import (
	"encoding/json"
	"net/http"
)

type UserController struct {
	httpServer *http.Server
	// ctx        context.Context
	// cancelFunc context.CancelFunc
	userManager *UserManager

	// orch
}

type RegisterUserRequest struct {
	Login    string
	Password string
}

type LogoutRequest struct {
	token string
}

func NewUserController(userManager *UserManager) *UserController {
	return &UserController{userManager: userManager}
}

func (ctrl *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var request RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = ctrl.userManager.RegisterUser(request.Login, request.Password)
	if err != nil {
		http.Error(w, "Failed to save user data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered successfully"))
}

func (ctrl *UserController) LoginApiRequest(w http.ResponseWriter, r *http.Request) {

	var request RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tokenStr, err := ctrl.userManager.LoginUser(request.Login, request.Password)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login credentials saved successfully. Token:" + tokenStr))
}

func (ctrl *UserController) LogoutApiRequest(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	tokenString, ok := requestData["token"]
	if !ok {
		http.Error(w, "Token not found in request", http.StatusBadRequest)
		return
	}

	err = ctrl.userManager.Logout(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
