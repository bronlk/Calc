package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestRegisterUserNewUser(t *testing.T) {
	um := NewUserManager()
	newUserLogin := "new"
	newUserPassword := "11"

	err := um.RegisterUser(newUserLogin, newUserPassword)
	if err != nil {
		t.Errorf("Expected no error for registering a new user, but got %v", err)
	}

	if _, exists := um.users[newUserLogin]; !exists {
		t.Errorf("Expected user %s to be registered, but they are not", newUserLogin)
	}
	um.RegisterUser("testuser", "testpassword")

	token, err := um.LoginUser("testuser", "testpassword")
	if err != nil {
		t.Errorf("Expected no error during login, got: %v", err)
	}

	if token == "" {
		t.Errorf("Expected token to be generated")
	}

	validUser, err := um.CheckJwt(token)
	if err != nil {
		t.Errorf("Expected no error checking JWT token, got: %v", err)
	}

	if validUser != "testuser" {
		t.Errorf("Expected JWT token to be valid for 'testuser'")
	}
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.a3gFh_Akbd9nHjmijXoxWvp-S9cQsKpfHl00rXKk6qk"

	login, err := um.CheckJwt(validToken)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedLogin := "testUser"
	if login != expectedLogin {
		t.Errorf("Expected login to be %s, but got %s", expectedLogin, login)
	}

}
func TestRegisterUserExistingUser(t *testing.T) {
	um := UserManager{}
	existingUserLogin := "old"
	existingUserPassword := "11"
	um.users[existingUserLogin] = "hashedPassword"

	err := um.RegisterUser(existingUserLogin, existingUserPassword)
	expectedErrorMessage := "User already exists"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message: '%s', but got: '%v'", expectedErrorMessage, err)
	}
}

func TestLoginUser(t *testing.T) {
	um := NewUserManager()

	um.RegisterUser("testuser", "testpassword")

	token, err := um.LoginUser("testuser", "testpassword")
	if err != nil {
		t.Errorf("Expected no error during login, got: %v", err)
	}

	if token == "" {
		t.Errorf("Expected token to be generated")
	}

	// validUser, err := um.CheckJwt(token)
	// if err != nil {
	// 	t.Errorf("Expected no error checking JWT token, got: %v", err)
	// }

	// if validUser != "testuser" {
	// 	t.Errorf("Expected JWT token to be valid for 'testuser'")
	// }
}

func TestCheckJwtValidToken(t *testing.T) {
	um := UserManager{}
	validToken := "VALID.JWT.TOKEN.STRING"

	login, err := um.CheckJwt(validToken)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedLogin := "testUser"
	if login != expectedLogin {
		t.Errorf("Expected login to be %s, but got %s", expectedLogin, login)
	}
}

func TestCheckJwtInvalidToken(t *testing.T) {
	um := UserManager{}
	invalidToken := "INVALID.JWT.TOKEN.STRING"

	_, err := um.CheckJwt(invalidToken)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	expectedErrorMessage := "Invalid token"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMessage, err.Error())
	}
}

type UserCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

const file string = "/home/ivan/prog/golang/golangmain/finaltests/mine/sprint2final_proj_ready/sqlite_db/sqlite_test.db"

func NewDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db := NewDBConnection()
	defer db.Close()

	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, login TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//data := UserCredentials{Login: user.Login}

	_, err = stmt.Exec()
	if err != nil {
		http.Error(w, "Failed to save login data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login credentials saved successfully"))
}

func PasswordHandler(w http.ResponseWriter, r *http.Request) {
	db := NewDBConnection()
	defer db.Close()

	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users(id TEXT, key TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Login, user.Password)
	if err != nil {
		http.Error(w, "Failed to save password data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password credentials saved successfully"))
}

func TestLoginHandler(t *testing.T) {

	// Регистрируем обработчики для каждого пути
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/password", PasswordHandler)

	// Запускаем сервер на порту :8080
	http.ListenAndServe(":8080", nil)
	fmt.Print("server is running")
}

// Формируем тело POST запроса с логином и паролем
//	body := []byte(`{"login": "testUser", "password": "testPassword"}`)

// Отправляем POST запрос на сервер
// resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(body))
// if err != nil {
// 	t.Fatalf("Error sending POST request: %v", err)
// }
// defer resp.Body.Close()

// Проверяем статус ответа сервера
//	if resp.StatusCode != http.StatusOK {
//		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
//}
