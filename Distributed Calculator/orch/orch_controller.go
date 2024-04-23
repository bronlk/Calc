package main

import (
	"encoding/json"
	"net/http"
)

type OrchController struct {
	httpServer *http.Server
	// ctx        context.Context
	// cancelFunc context.CancelFunc
	orch        *Orchestrator
	userManager *UserManager

	// orch
}

type ExpressionRequest struct {
	Expression string
	Token      string
}

type ExpressionCalcRequest struct {
	CalcId string
}

func NewOrchestratorController(orch *Orchestrator, userManager *UserManager) *OrchController {
	return &OrchController{orch: orch, userManager: userManager}
}

// func GetExpressionByApi(userManager *UserManager) *UserController {
// 	return &UserController{userManager: userManager}
// }

func (c *OrchController) AddExpressionByApi(w http.ResponseWriter, r *http.Request) {

	var request ExpressionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	login, err := c.userManager.checkJwt(request.Token)

	if err != nil {
		http.Error(w, "Вы не залогинены", http.StatusUnauthorized)
		return
	}

	var expr = Expression{Id: -1, Login: login, Expression: request.Expression, Result: "", Status: "New"}
	c.orch.Add(expr)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Expression saved successfully"))
}

// вызывается клькулятором
func (orchController *OrchController) GetExpressionByApi(w http.ResponseWriter, r *http.Request) {
	var request ExpressionCalcRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	expression, err := orchController.orch.GetExpressionForCalc(request.CalcId)

	if err != nil {
		http.Error(w, "Expression not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Your expression: " + expression.Expression))
}

func (orchController *OrchController) PrintExpressionByApi(w http.ResponseWriter, r *http.Request) {

	var request ExpressionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	orchController.orch.PrintExpressions()

	w.WriteHeader(http.StatusOK)
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
