package main

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/laurent22/go-sqlkv"
	_ "github.com/mattn/go-sqlite3"
)

type UserManager struct {
	//httpServer *http.Server
	userRepo *UserRepository

	// orch
}

func NewUserManager(userRepo *UserRepository) *UserManager {
	return &UserManager{
		userRepo: userRepo,
	}
}

func (userManager *UserManager) RegisterUser(login, password string) error {

	user := User{
		Login:    login,
		Password: string(password),
		Active:   true,
	}

	err := userManager.userRepo.InsertUser(user)

	return err
}

func (userManager *UserManager) LoginUser(login, password string) (string, error) {
	user, err := userManager.userRepo.SelectUserByLogin(login)

	if user.Password != password {
		return "", errors.New("Invalid password")
	}

	claims := jwt.MapClaims{
		"login": login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte("0")
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	err = userManager.userRepo.UserTokenSet(user.ID, tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (userManager *UserManager) Logout(tokenString string) error {

	login, err := userManager.checkJwt(tokenString)

	userManager.userRepo.UserTokenClear(login)
	return err
}

func (userManager *UserManager) checkJwt(tokenString string) (string, error) {
	key := []byte("0")
	// parse jwt token
	// get login
	// userRepo.UserTokenGet(login)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}
	// Проверяем, что токен валиден
	if !token.Valid {
		return "", fmt.Errorf("Token is not valid")
	}
	// Извлекаем login из содержимого токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Failed to parse claims")
	}
	login, ok := claims["login"].(string)
	if !ok {
		return "", fmt.Errorf("Login not found in claims")
	}
	return login, nil
}

// token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 	return []byte("0"), nil
// })
// if err != nil {
// 	return "", err
// }
// if token.Valid {
// 	claims := token.Claims.(jwt.MapClaims)
// 	login := claims["login"].(string)
// 	// Проверка токена с базой данных
// 	var savedLogin string
// 	err := db.QueryRow("SELECT login FROM users WHERE login = ?", login).Scan(&savedLogin)
// 	if err != nil {
// 		return "", errors.New("User not found")
// 	}
// 	if savedLogin != login {
// 		return "", errors.New("Invalid login")
// 	}
// 	return login, nil
// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// if err != nil {
// 	return errors.New("Invalid password")
// }
// key := make([]byte, 32)
// _, err = rand.Read(key)
// if err != nil {
// 	return err
// }
// token := jwt.New(jwt.SigningMethodHS256)
// tokenString, err := token.SignedString(key)
// if err != nil {
// 	return err
// }

// func LoginUser(db *sql.DB, login, password string) error {
// 	var user User
// 	row := db.QueryRow("SELECT id, login, password, active, jwtkey FROM users WHERE login = ?", login)
// 	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.Active)
// 	if err != nil {
// 		return errors.New("Invalid login")
// 	}
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// 	if err != nil {
// 		return errors.New("Invalid password")
// 	}
// 	key := []byte("mhgfjhytfcjyrtdjyrcjytrshtrshtrhtsd")
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	tokenString, err := token.SignedString(key)
// }

// func FindLoginFromToken(tokenString string, key []byte) (string, error) {
// 	// Парсим токен
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return key, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	// Проверяем, что токен валиден
// 	if !token.Valid {
// 		return "", fmt.Errorf("Token is not valid")
// 	}
// 	// Извлекаем login из содержимого токена
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return "", fmt.Errorf("Failed to parse claims")
// 	}
// 	login, ok := claims["login"].(string)
// 	if !ok {
// 		return "", fmt.Errorf("Login not found in claims")
// 	}
// 	return login, nil
// }
// func (userManager *UserManager) JwtHandler(tokenString string) string {
// 	key := []byte("0") // Ключ должен совпадать с ключом, использованным при создании токена
// 	login, err := FindLoginFromToken(tokenString, key)
// 	if err != nil {
// 		fmt.Printf("Failed to find login: %v\n", err)
// 		return
// 	}
// 	fmt.Println("Login found in token:", login)
// }
