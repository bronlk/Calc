package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Expression struct {
	Id         string
	Expression string
	Result     string
	Status     string
}

type ExpressionAgent struct {
	clientId string
	working  bool
	wg       sync.WaitGroup
}

func NewExpressionAgent(clientId string) *ExpressionAgent {
	return &ExpressionAgent{
		clientId: clientId,
		working:  false,
	}
}

func (ea *ExpressionAgent) Start() {
	ea.working = true
	go ea.mainLoop()
	ea.wg.Add(1)
}

func (ea *ExpressionAgent) StopTimeout(timeout time.Duration) error {
	c := make(chan struct{})
	go func() {
		defer close(c)
		ea.wg.Wait()
	}()
	select {
	case <-c:
		return nil
	case <-time.After(timeout):
		return errors.New("exit by timeout")
	}
}

func (ea *ExpressionAgent) Stop() {
	ea.working = false
	ea.wg.Wait()
}

func (ea *ExpressionAgent) mainLoop() {
	defer ea.wg.Done()

	// Подключение к базе данных
	db, err := sql.Open("sqlite3", "../sqlite_db/sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for ea.working {
		// Обновление last_connect для агента
		_, err = db.Exec("UPDATE agents SET last_connect = ? WHERE name = ?", time.Now(), ea.clientId)
		if err != nil {
			fmt.Println("Ошибка при обновлении last_connect:", err)
			continue // Пропускаем итерацию при ошибке
		}

		// Запрос на получение необработанного выражения
		row := db.QueryRow("SELECT id, expression FROM expressions WHERE result IS '' LIMIT 1")

		var id int
		var expression string
		err = row.Scan(&id, &expression)
		if err == sql.ErrNoRows {
			// Нет необработанных выражений, ждем 2 секунды
			time.Sleep(2 * time.Second)
			continue
		} else if err != nil {
			fmt.Println("Ошибка при запросе:", err)
			continue // Пропускаем итерацию при ошибке
		}

		// Вычисление выражения
		result, err := calculate(expression)
		if err != nil {
			fmt.Println("Ошибка вычисления:", err)
			continue // Пропускаем итерацию при ошибке
		}

		// Обновление записи в базе данных (expressions и agents)
		tx, err := db.Begin() // Начинаем транзакцию
		if err != nil {
			fmt.Println("Ошибка при начале транзакции:", err)
			continue // Пропускаем итерацию при ошибке
		}
		_, err = tx.Exec("UPDATE expressions SET result = ? WHERE id = ?", result, id)
		if err != nil {
			fmt.Println("Ошибка при обновлении выражения:", err)
			tx.Rollback() // Откатываем транзакцию при ошибке
			continue      // Пропускаем итерацию при ошибке
		}
		_, err = tx.Exec("UPDATE agents SET expr_id = ? WHERE name = ?", id, ea.clientId)
		if err != nil {
			fmt.Println("Ошибка при обновлении агента:", err)
			tx.Rollback() // Откатываем транзакцию при ошибке
			continue      // Пропускаем итерацию при ошибке
		}
		err = tx.Commit() // Коммитим транзакцию
		if err != nil {
			fmt.Println("Ошибка при коммите транзакции:", err)
			continue
		}

		fmt.Println("Выражение вычислено:", expression, "=", result)
	}
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func calculate(equation string) (int64, error) {

	var operator string
	if strings.Contains(equation, "+") {
		operator = "+"
	} else if strings.Contains(equation, "-") {
		operator = "-"
	} else if strings.Contains(equation, "*") {
		operator = "*"
	} else if strings.Contains(equation, ":") {
		operator = ":"
	} else {
		return 0, fmt.Errorf("Invalid equation")
	}

	parts := strings.Split(equation, operator)

	operand1, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	operand2, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, err
	}

	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case ":":
		if operand2 == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return operand1 / operand2, nil
	default:
		return 0, fmt.Errorf("Invalid equation")
	}
}

func doCalc(id string, exp Expression, httpCl *http.Client) {
	res, _ := calculate(exp.Expression)
	exp.Result = strconv.Itoa(int(res))
	exp.Status = id
	expBytes, _ := json.Marshal(exp)

	httpCl.Post("http://localhost:8080/set_expression_result", "application/json", bytes.NewBuffer(expBytes))
}

func main() {
	name := "Calculator" // или name := os.Args[1] для аргумента командной строки

	fmt.Println("Starting calc name:", name)

	// Подключение к базе данных
	db, err := sql.Open("sqlite3", "../sqlite_db/sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Создание таблиц, если они не существуют

	// Добавление агента в базу данных
	_, err = db.Exec("INSERT INTO agents (name, last_connect) VALUES (?, ?)", name, time.Now())
	if err != nil {
		fmt.Println("Ошибка при добавлении агента:", err)
	}

	var agent *ExpressionAgent = NewExpressionAgent(name)
	agent.Start()

	fmt.Printf("Press Enter to stop")
	fmt.Fscanln(os.Stdin)
	fmt.Println("Stopping")

	agent.Stop()
	fmt.Println("Stopped")
}
