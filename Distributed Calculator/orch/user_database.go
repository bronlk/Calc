package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Login    string
	Password string
	Active   bool
}
type UserRepository struct {
	fileName string
}

func NewUserRepository(fileName string) *UserRepository {
	return &UserRepository{fileName: fileName}
}

func InitDB(fileName string) error {
	db, err := openDbConnection(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	users := `
	CREATE TABLE IF NOT EXISTS users (
		"id"			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"login"			TEXT NOT NULL UNIQUE,
		"password"		TEXT,
		"jwtkey"		TEXT,
		"active"		INTEGER		
	);
`
	expressions := `
	CREATE TABLE IF NOT EXISTS expressions (
		id TEXT PRIMARY KEY UNIQUE,
		expression TEXT,
		result TEXT,
		status TEXT
	);
`

	db.Exec(users, expressions)
	return nil
}

func openDbConnection(fileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func (userRepo *UserRepository) InsertUser(user User) error {
	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(login, password, active) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Login, user.Password, user.Active)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (userRepo *UserRepository) SelectUserByLogin(login string) (User, error) {
	var user User

	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()
	row := db.QueryRow("SELECT id, login, password, active FROM users WHERE login = ?", login)
	err = row.Scan(&user.ID, &user.Login, &user.Password, &user.Active)
	if err != nil {
		return User{}, err
	}
	db.Close()
	return user, nil
}

func (userRepo *UserRepository) SelectUserById(db *sql.DB, userID int) (User, error) {
	var user User

	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()
	row := db.QueryRow("SELECT id, login, password, active FROM users WHERE id = ?", userID)
	err = row.Scan(&user.ID, &user.Login, &user.Password, &user.Active)
	if err != nil {
		return User{}, err
	}
	db.Close()
	return user, nil
}

// func (userRepo *UserRepository) UserTokenGet(login string) error {
// 	db, err := openDbConnection(userRepo.fileName)
// 	defer db.Close()
// 	//_, err = db.Exec("UPDATE users SET jwtkey = '' WHERE login = ?", login)
// 	if err != nil {
// 		return err
// 	}
// 	db.Close()
// 	return nil
// }

func (userRepo *UserRepository) UserTokenSet(userID int, tokenString string) error {
	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()
	_, err = db.Exec("UPDATE users SET jwtkey = ? WHERE id = ?", tokenString, &userID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (userRepo *UserRepository) UserTokenClear(login string) error {
	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()
	_, err = db.Exec("UPDATE users SET jwtkey = '' WHERE login = ?", login)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

// func (userRepo *UserRepository) User(tokenString string) error {
// 	db, err := openDbConnection(userRepo.fileName)
// 	defer db.Close()
// 	login, err := userManager.JwtHandler(tokenString)
// 	if err != nil {
// 		return err
// 	}
// 	// Получаем ID пользователя по логину
// 	var userID int
// 	err = db.QueryRow("SELECT id FROM users WHERE login = ?", login).Scan(&userID)
// 	if err != nil {
// 		return err
// 	}
// 	err = userRepo.deactivateUser(db, userID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (userRepo *UserRepository) DeleteUser(db *sql.DB, userID int) error {
	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()
	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (userRepo *UserRepository) DeactivateUser(db *sql.DB, userID int) error {
	db, err := openDbConnection(userRepo.fileName)
	defer db.Close()
	_, err = db.Exec("UPDATE users SET active = 0 WHERE id = ?", userID)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}

func (orch *Orchestrator) SaveExpression(string) error {
	db, _ := openDbConnection("../sqlite_db/sqlite.db")
	defer db.Close()
	for _, exp := range orch.Expressions() {
		_, err := db.Exec("INSERT INTO expressions (id, expression, result, status) VALUES (?, ?, ?, ?)", exp.Id, exp.Expression, exp.Result, exp.Status)
		if err != nil {
			return err
		}
	}

	return nil
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

func (orch *Orchestrator) AddExpression(expressionText string) string {
	var id string
	err := orch.db.QueryRow("INSERT INTO expressions(expression, result, status) VALUES(?, '', 'New') RETURNING id", expressionText).Scan(&id)
	if err != nil {
		fmt.Println("Error inserting new expression into database:", err)
		return ""
	}

	_, err = orch.db.Exec("UPDATE agents SET expr_id = ?, last_connect = ? WHERE name = ?", id, time.Now())
	if err != nil {
		fmt.Println("Error assigning expression to agent in database:", err)
		return ""
	}

	return id
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
