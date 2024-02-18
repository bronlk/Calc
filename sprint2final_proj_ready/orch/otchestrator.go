package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/zeroflucs-given/generics/collections/stack"
)

type OrchestratorCfg struct {
	Expressions []Expression
	IdCounter   int
	Queue       []Expression
}

type Orchestrator struct {
	mapSync       *sync.Mutex
	expressionMap map[string]*Expression
	agentMap      map[string]*CalcAgent
	idCounter     int
	stack         *stack.Stack[*Expression]
}

type CalcAgent struct {
	Name        string
	ExprId      string
	lastConnect time.Time
}

type Expression struct {
	Id         string
	Expression string
	Result     string
	Status     string
}

func NewOrchestrator() *Orchestrator {
	m := make(map[string]*Expression)
	agents := make(map[string]*CalcAgent)
	var lock = &sync.Mutex{}
	var st = stack.NewStack[*Expression](1000)
	return &Orchestrator{expressionMap: m, mapSync: lock, stack: st, agentMap: agents}
}

func (orch *Orchestrator) ApplyCfg(cfg OrchestratorCfg) {

	orch.idCounter = cfg.IdCounter

	for _, expr := range cfg.Expressions {
		orch.expressionMap[expr.Id] = &Expression{Id: expr.Id,
			Expression: expr.Expression, Result: expr.Result, Status: expr.Status}
	}

	for _, expr := range cfg.Queue {
		orch.stack.Push(&expr)
	}
}

func (orch *Orchestrator) GetCfg() OrchestratorCfg {

	exprs := make([]Expression, 0, len(orch.expressionMap))

	for _, value := range orch.expressionMap {
		exprs = append(exprs,
			Expression{Id: value.Id,
				Expression: value.Expression, Result: value.Result, Status: value.Status})
	}

	queue := make([]Expression, 0, orch.stack.Count())

	// for _, vl := range orch.stack.Pop() {
	// 	queue = append(queue, vl.(Expression))
	// }

	var res OrchestratorCfg = OrchestratorCfg{IdCounter: orch.idCounter, Expressions: exprs, Queue: queue}

	return res
}

func (orch *Orchestrator) Add(expressionText string) string {

	orch.mapSync.Lock()

	idStr := strconv.Itoa(orch.idCounter)
	var exp Expression = Expression{Id: idStr, Expression: expressionText, Result: "", Status: "New"}
	orch.idCounter++
	orch.expressionMap[idStr] = &exp
	orch.stack.Push(&exp)
	//fmt.Println("Add" + strconv.Itoa(len(*orch.queue)))

	orch.mapSync.Unlock()

	return idStr
}

// список выражений со статусом вычислений
func (orch *Orchestrator) Expressions() []Expression {

	orch.mapSync.Lock()

	orch.CheckStatus()

	exprArr := make([]Expression, len(orch.expressionMap))
	var idx int = 0
	for _, value := range orch.expressionMap {
		exprArr[idx] = Expression{Id: value.Id,
			Expression: value.Expression, Result: value.Result, Status: value.Status}
		idx++
	}

	orch.mapSync.Unlock()

	return exprArr
}

// список выражений со статусом вычислений
func (orch *Orchestrator) Agents() []CalcAgent {

	orch.mapSync.Lock()

	orch.CheckStatus()

	exprArr := make([]CalcAgent, len(orch.agentMap))
	var idx int = 0
	for _, value := range orch.agentMap {
		exprArr[idx] = CalcAgent{Name: value.Name, ExprId: value.ExprId, lastConnect: value.lastConnect}
		idx++
	}

	orch.mapSync.Unlock()

	return exprArr
}

func (orch *Orchestrator) CheckStatus() {

	time := time.Now()

	for _, agent := range orch.agentMap {

		ff := time.Sub(agent.lastConnect)

		if ff.Abs().Milliseconds() > 2000 {

			expr, found := orch.expressionMap[agent.ExprId]

			if found {
				orch.stack.Push(expr)
				//fmt.Println("CheckStatus" + strconv.Itoa(len(*orch.queue)))
				expr.Status = "New"
			}
			delete(orch.agentMap, agent.Name)
		}
	}
}

func (orch *Orchestrator) updateAgentTime(agentId string) {

	agent, found := orch.agentMap[agentId]

	if !found {
		agent = &CalcAgent{Name: agentId, ExprId: ""}
		orch.agentMap[agentId] = agent
	}

	agent.lastConnect = time.Now()
}

func (orch *Orchestrator) GetExpression(calcId string) (bool, *Expression) {

	orch.mapSync.Lock()
	defer orch.mapSync.Unlock()

	agent, found := orch.agentMap[calcId]

	if !found {
		agent = &CalcAgent{Name: calcId, ExprId: ""}
		orch.agentMap[calcId] = agent
	}

	orch.agentMap[calcId].lastConnect = time.Now()

	found, expr := orch.stack.Pop()
	if found {
		//fmt.Println("GetExpression" + strconv.Itoa(len(*orch.queue)))

		orch.agentMap[calcId].ExprId = expr.Id
		orch.expressionMap[expr.Id].Status = "In calc:" + calcId

		fmt.Println("returned expr.Id:" + expr.Id)
		agent.lastConnect = time.Now()

		return true, expr
	} else {
		agent.lastConnect = time.Now()
		orch.agentMap[calcId] = agent
		return false, nil
	}

}

func (orch *Orchestrator) SetResult(exp Expression) {
	orch.mapSync.Lock()

	orch.updateAgentTime(exp.Status)

	expr, found := orch.expressionMap[exp.Id]

	if !found {
		fmt.Println("Expression with id" + exp.Id + " not found")
		return
	}
	expr.Result = exp.Result
	expr.Status = "done"

	orch.mapSync.Unlock()
}
