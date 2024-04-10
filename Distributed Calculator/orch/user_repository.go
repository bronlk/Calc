package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "../sqlite_db/sqlite.db"

type User struct {
	ID       int
	Login    string
	Password string
	Active   bool
}

func InitDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	query := `
	CREATE TABLE IF NOT EXISTS users (
		login TEXT,
		password TEXT,
		active INTEGER,
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE
	)
`
	db.Exec(query)
	return db
}

func insertUser(db *sql.DB, user User) error {
	stmt, err := db.Prepare("INSERT INTO users(login, password, active) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Login, user.Password, user.Active)
	if err != nil {
		return err
	}

	return nil
}

func selectUser(db *sql.DB, userID int) (User, error) {
	var user User

	row := db.QueryRow("SELECT id, login, password, active FROM users WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.Active)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func deleteUser(db *sql.DB, userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}

func deactivateUser(db *sql.DB, userID int) error {
	_, err := db.Exec("UPDATE users SET active = 0 WHERE id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}
