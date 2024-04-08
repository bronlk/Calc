package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/laurent22/go-sqlkv"
	_ "github.com/laurent22/go-sqlkv"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	users  map[string]string     // Мап для хранения пар login:захешированный_пароль
	tokens map[*jwt.Token]string // Мап для хранения токенов пользователей
}

type JWTKeyManager struct {
	DB *sql.DB
}

func NewUserManager() *UserManager {
	return &UserManager{
		users:  make(map[string]string),
		tokens: make(map[*jwt.Token]string),
	}
}

func (um *UserManager) RegisterUser(login, password string) error {

	_, exists := um.users[login]
	if exists {
		return errors.New("User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	um.users[login] = string(hashedPassword)
	return nil
}

func (um *UserManager) LoginUser(login, password string) (string, error) {
	const file string = "/home/ivan/prog/golang/golangmain/finaltests/mine/sprint2final_proj_ready/sqlite_db/sqlite.db"
	db, err := sql.Open("sqlite3", file)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	hashedPassword, ok := um.users[login]
	if !ok {
		return "", errors.New("Invalid login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("Invalid password")
	}

	key := []byte("mhgfjhytfcjyrtdjyrcjytrshtrshtrhtsd")
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(key)

	store := sqlkv.New(db, "jwtkey")
	store.SetString(login, tokenString)

	for token := range um.tokens {
		parsedToken, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return token, nil
		})

		if parsedToken.Valid {
			continue
		}

		delete(um.tokens, token)
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["login"] = login
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	if err != nil {
		return "", err
	}

	um.tokens[token] = login
	return tokenString, nil
}

func (um *UserManager) CheckJwt(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
		return tkn, nil
	})
	if err != nil {
		return "", err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		login := claims["login"].(string)
		return login, nil
	}

	return "", errors.New("Invalid token")
}

func (um *UserManager) Logout(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
		return tkn, nil
	})
	if err != nil {
		return err
	}

	if token.Valid {
		//claims := token.Claims.(jwt.MapClaims)
		//login := claims["login"].(string)
		delete(um.tokens, token)
		return nil
	}

	return errors.New("Invalid token")
}
