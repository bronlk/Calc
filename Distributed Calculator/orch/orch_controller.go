package main

import (
	"encoding/json"
	"net/http"
)

type OrchController struct {
	httpServer *http.Server
	// ctx        context.Context
	// cancelFunc context.CancelFunc
	orch *Orchestrator

	// orch
}

type ExpressionRequest struct {
	Expression string
}

func NewOrchestatorController(orch *Orchestrator) *OrchController {
	return &OrchController{orch: orch}
}

// func GetExpressionByApi(userManager *UserManager) *UserController {
// 	return &UserController{userManager: userManager}
// }

func (OrchController *OrchController) AddExpressionByApi(w http.ResponseWriter, r *http.Request) {

	var request ExpressionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = orch.addExpression(request.Expression)
	if err != nil {
		http.Error(w, "Failed to save login data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Registration credentials saved successfully"))
}

// func (userController *UserController) LoginApiRequest(w http.ResponseWriter, r *http.Request) {

// 	var request RegisterUserRequest
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	err = userController.userManager.LoginUser(request.Login, request.Password)
// 	if err != nil {
// 		http.Error(w, "Failed to login", http.StatusUnauthorized)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Login credentials saved successfully"))
// }

// func (userController *UserController) LogoutApiRequest(w http.ResponseWriter, r *http.Request) {
// 	var requestData map[string]string
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
// 		return
// 	}

// 	tokenString, ok := requestData["token"]
// 	if !ok {
// 		http.Error(w, "Token not found in request", http.StatusBadRequest)
// 		return
// 	}

// 	err = userController.userManager.Logout(tokenString)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Logout successful"))
// }
