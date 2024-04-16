package main

import (
	"database/sql"
	"log"

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
	query := `
	CREATE TABLE IF NOT EXISTS users (
		"id"			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
		"login"			TEXT NOT NULL UNIQUE,
		"password"		TEXT,
		"jwtkey"		TEXT,
		"active"		INTEGER		
		);
`
	db.Exec(query)
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
