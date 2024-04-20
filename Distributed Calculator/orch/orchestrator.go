package main

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Orchestrator struct {
	db       *sql.DB
	mapSync  *sync.Mutex
	orchRepo *OrchRepository
}

type CalcAgent struct {
	Name        string
	ExprId      string
	LastConnect time.Time
}

type Expression struct {
	Id         int
	Login      string
	Expression string
	Result     string
	Status     string
}

func NewOrchestrator(orchRepo *OrchRepository) *Orchestrator {
	return &Orchestrator{
		orchRepo: orchRepo,
	}
}

func (orch *Orchestrator) Add(expression Expression) {
	orch.orchRepo.SaveExpression(expression)

	fmt.Print("Saved successfully")
}

func (orch *Orchestrator) GetExpressionForCalc(calcId string) (*Expression, error) {
	exp, err := orch.orchRepo.ObtainExpressionForCalc(calcId)
	if err !=nil {
		return nil, err
	}
	return exp, nil
}

func (orch *Orchestrator) SetResult(exp Expression) string {
	err := orch.SetResultByID(&exp)
	return err
}

// func (orch *Orchestrator) Add(expressionText string) string {
// 	err := orch.AddExpression(expressionText)
// 	return err
// }

func (orch *Orchestrator) Expressions() []Expression {
	err := orch.PrintExpressions()
	return err
}
