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
	Id         string
	Expression string
	Result     string
	Status     string
}

func NewOrchestrator(orchRepo *OrchRepository) *Orchestrator {
	return &Orchestrator{
		orchRepo: orchRepo,
	}
}

func (orch *Orchestrator) SaveToDatabase(expressionText string) {
	err := orch.SaveExpression(expressionText)
	if err != nil {
		fmt.Printf("Failed to save")
		return
	}
	fmt.Print("Saved successfully")
}

func (orch *Orchestrator) GetExpression(calcId string) {
	err, _ := orch.GetExpressionByAgent(calcId)
	if err != false {
		fmt.Print("Saved successfully")
		return
	}
	fmt.Print("Saved successfully")
}

func (orch *Orchestrator) SetResult(exp Expression) string {
	err := orch.SetResultByID(&exp)
	return err
}

func (orch *Orchestrator) Add(expressionText string) string {
	err := orch.AddExpression(expressionText)
	return err
}

func (orch *Orchestrator) Expressions() []Expression {
	err := orch.PrintExpressions()
	return err
}
