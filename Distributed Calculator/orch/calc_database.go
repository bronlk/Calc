package main

import (
	"context"
	"fmt"
	"log"
	"net/http" // Путь к вашим структурам
	"time"
)

type OrchRepository struct {
	httpServer *http.Server
	ctx        context.Context
	cancelFunc context.CancelFunc
	orch       *Orchestrator
	fileName   string // Используйте вашу структуру Orchestrator
}

func NewOrchRepository(fileName string) *OrchRepository {
	return &OrchRepository{fileName: fileName}
}

func InitOrchDB(fileName string) error {
	db, err := openDbConnection(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	expressions := `
	CREATE TABLE IF NOT EXISTS expressions (
		id TEXT PRIMARY KEY UNIQUE,
		expression TEXT,
		result TEXT,
		status TEXT
	);
`

	db.Exec(expressions)
	return nil
}

// func (orch *Orchestrator) SaveExpression(string) error {
// 	db, _ := openDbConnection("../sqlite_db/sqlite.db")
// 	defer db.Close()
// 	exp := orch.Expressions()
// 	_, err := db.Exec("INSERT INTO expressions (id, expression, result, status) VALUES (?, ?, ?, ?)", exp.Id, exp.Expression, exp.Result, exp.Status)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (orch *Orchestrator) SaveExpression(expression string) string {
	db, err := openDbConnection(orch.orchRepo.fileName)
	defer db.Close()
	var exp *Expression
	err = orch.db.QueryRow("INSERT INTO expressions(expression, result, status) VALUES(?, '', 'New') RETURNING id", expression).Scan(&exp.Id)
	if err != nil {
		fmt.Println("Error inserting new expression into database:", err)
		return ""
	}

	_, err = orch.db.Exec("UPDATE agents SET expr_id = ?, last_connect = ? WHERE name = ?", exp.Id, time.Now())
	if err != nil {
		fmt.Println("Error assigning expression to agent in database:", err)
		return ""
	}

	return exp.Id
}

func (orch *Orchestrator) GetExpressionByAgent(calcId string) (bool, *Expression) {
	var expr Expression
	var agent CalcAgent

	row := orch.db.QueryRow("SELECT expr_id, status FROM agents WHERE name = ?", calcId)
	err := row.Scan(&agent.ExprId, &expr.Status)
	if err != nil {
		return false, nil
	}

	row = orch.db.QueryRow("SELECT id, expression, result FROM expressions WHERE id = ?", agent.ExprId)
	err = row.Scan(&expr.Id, &expr.Expression, &expr.Result)
	if err != nil {
		return false, nil
	}

	_, err = orch.db.Exec("UPDATE expressions SET status = ? WHERE id = ?", "In calc:"+calcId, expr.Id)
	if err != nil {
		fmt.Println("Error updating expression status in database:", err)
	}

	_, err = orch.db.Exec("UPDATE agents SET last_connect = ? WHERE name = ?", time.Now(), calcId)
	if err != nil {
		fmt.Println("Error updating agent last connect time in database:", err)
	}

	return true, &expr
}

func (orch *Orchestrator) SetResultByID(exp *Expression) string {
	_, err := orch.db.Exec("UPDATE expressions SET result = ?, status = 'done' WHERE id = ?", exp.Result, exp.Id)
	if err != nil {
		fmt.Println("Error updating expression result in database:", err)
	}
	return "Result updated successfully"
}

func (orch *Orchestrator) PrintExpressions() []Expression {
	rows, err := orch.db.Query("SELECT id, expression, result, status FROM expressions")
	if err != nil {
		fmt.Println("Error fetching expressions from database:", err)
		return nil
	}
	defer rows.Close()

	var expressions []Expression
	for rows.Next() {
		var expr Expression
		err := rows.Scan(&expr.Id, &expr.Expression, &expr.Result, &expr.Status)
		if err != nil {
			fmt.Println("Error scanning expression:", err)
			continue
		}
		expressions = append(expressions, expr)
	}

	return expressions
}
