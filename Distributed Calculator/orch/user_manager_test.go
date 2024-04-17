package main

// func TestRegisterUserNewUser(t *testing.T) {
// 	//	um := NewUserManager()
// 	newUserLogin := "new"
// 	newUserPassword := "11"

// 	err := um.RegisterUser(newUserLogin, newUserPassword)
// 	if err != nil {
// 		t.Errorf("Expected no error for registering a new user, but got %v", err)
// 	}

// 	if _, exists := um.users[newUserLogin]; !exists {
// 		t.Errorf("Expected user %s to be registered, but they are not", newUserLogin)
// 	}
// 	um.RegisterUser("testuser", "testpassword")

// 	token, err := um.LoginUser("testuser", "testpassword")
// 	if err != nil {
// 		t.Errorf("Expected no error during login, got: %v", err)
// 	}

// 	if token == "" {
// 		t.Errorf("Expected token to be generated")
// 	}

// 	validUser, err := um.CheckJwt(token)
// 	if err != nil {
// 		t.Errorf("Expected no error checking JWT token, got: %v", err)
// 	}

// 	if validUser != "testuser" {
// 		t.Errorf("Expected JWT token to be valid for 'testuser'")
// 	}
// 	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.a3gFh_Akbd9nHjmijXoxWvp-S9cQsKpfHl00rXKk6qk"

// 	login, err := um.CheckJwt(validToken)
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	expectedLogin := "testUser"
// 	if login != expectedLogin {
// 		t.Errorf("Expected login to be %s, but got %s", expectedLogin, login)
// 	}

// }
// func TestRegisterUserExistingUser(t *testing.T) {
// 	um := UserManager{}
// 	existingUserLogin := "old"
// 	existingUserPassword := "11"
// 	um.users[existingUserLogin] = "hashedPassword"

// 	err := um.RegisterUser(existingUserLogin, existingUserPassword)
// 	expectedErrorMessage := "User already exists"
// 	if err == nil || err.Error() != expectedErrorMessage {
// 		t.Errorf("Expected error message: '%s', but got: '%v'", expectedErrorMessage, err)
// 	}
// }

// func TestLoginUser(t *testing.T) {
// 	um := NewUserManager()

// 	um.RegisterUser("testuser", "testpassword")

// 	token, err := um.LoginUser("testuser", "testpassword")
// 	if err != nil {
// 		t.Errorf("Expected no error during login, got: %v", err)
// 	}

// 	if token == "" {
// 		t.Errorf("Expected token to be generated")
// 	}

// validUser, err := um.CheckJwt(token)
// if err != nil {
// 	t.Errorf("Expected no error checking JWT token, got: %v", err)
// }

// if validUser != "testuser" {
// 	t.Errorf("Expected JWT token to be valid for 'testuser'")
// }
//}

// func TestCheckJwtValidToken(t *testing.T) {
// 	um := UserManager{}
// 	validToken := "VALID.JWT.TOKEN.STRING"

// 	login, err := um.CheckJwt(validToken)
// 	if err != nil {
// 		t.Errorf("Expected no error, got %v", err)
// 	}

// 	expectedLogin := "testUser"
// 	if login != expectedLogin {
// 		t.Errorf("Expected login to be %s, but got %s", expectedLogin, login)
// 	}
// }

// func TestCheckJwtInvalidToken(t *testing.T) {
// 	um := UserManager{}
// 	invalidToken := "INVALID.JWT.TOKEN.STRING"

// 	_, err := um.CheckJwt(invalidToken)
// 	if err == nil {
// 		t.Error("Expected an error, but got nil")
// 	}

// 	expectedErrorMessage := "Invalid token"
// 	if err.Error() != expectedErrorMessage {
// 		t.Errorf("Expected error message: '%s', but got: '%s'", expectedErrorMessage, err.Error())
// 	}
// }

// type UserCredentials struct {
// 	Login    string `json:"login"`
// 	Password string `json:"password"`
// }

//const file string = "/home/ivan/prog/golang/golangmain/finaltests/mine/sprint2final_proj_ready/sqlite_db/sqlite.db"

// func NewDBConnection() *sql.DB {
// 	db, err := sql.Open("sqlite3", file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		login TEXT,
// 		password TEXT,
// 		active INTEGER,
// 		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE
// 	)
// `
// 	db.Exec(query)
// 	return db
// }

// func RegApiRequests(w http.ResponseWriter, r *http.Request) {
// 	db := NewDBConnection()
// 	defer db.Close()
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	err = RegisterUser(db, user.Login, user.Password)
// 	if err != nil {
// 		http.Error(w, "Failed to save login data", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Registration credentials saved successfully"))
// }

// func LoginApiRequests(w http.ResponseWriter, r *http.Request) {
// 	db := NewDBConnection()
// 	defer db.Close()
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	err = LoginUser(db, user.Login, user.Password)
// 	if err != nil {
// 		http.Error(w, "Failed to login", http.StatusUnauthorized)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Login credentials saved successfully"))
// }

// func LogoutApiRequests(w http.ResponseWriter, r *http.Request) {
// 	db := NewDBConnection()
// 	defer db.Close()

// 	var user User
// 	err := Logout(db, user.Login)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Logout successful"))
// }

// func TestLoginHandler(t *testing.T) {

// 	// Регистрируем обработчики для каждого пути
// 	http.HandleFunc("/reg", RegApiRequests)
// 	http.HandleFunc("/login", LoginApiRequests)
// 	http.HandleFunc("/logout", LogoutApiRequests)
// 	// http.HandleFunc("/password", PasswordHandler)

// 	// Запускаем сервер на порту :8080
// 	http.ListenAndServe(":8080", nil)
// 	fmt.Print("server is running")
// }

// Формируем тело POST запроса с логином и паролем
// 	body := []byte(`{"login": "testUser", "password": "testPassword"}`)

// 		// Отправляем POST запрос на сервер
// resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(body))
// if err != nil {
// 	t.Fatalf("Error sending POST request: %v", err)
// }
// defer resp.Body.Close()

// Проверяем статус ответа сервера
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
// }

// func TestUserRepo(t *testing.T) {
// 	db := InitDBConnection()
// 	defer db.Close()
// 	// Пример использования функций
// 	newUser := User{Login: "testUser", Password: "testPassword", Active: true}
// 	err := insertUser(db, newUser)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	selectedUser, err := selectUser(db, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Selected User:", selectedUser)

// 	err = deleteUser(db, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = deactivateUser(db, 2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func TestUserManagement(t *testing.T) {
// 	db := InitDBConnection()
// 	defer db.Close()
// 	logindata := "testU12ser121411133123"
// 	err := RegisterUser(db, logindata, "testPassword")
// 	if err != nil {
// 		t.Errorf("RegisterUser failed. Error: %v", err)
// 	}

// 	err = LoginUser(db, logindata, "testPassword")
// 	if err != nil {
// 		t.Errorf("LoginUser failed. Error: %v", err)
// 	}

// 	token := "your-generated-jwt-token"
// 	login, err := CheckJwt(db, token)
// 	if err != nil {
// 		t.Errorf("CheckJwt failed. Error: %v", login)
// 	}

// 	err = Logout(db, token)
// 	if err != nil {
// 		t.Errorf("Logout failed. Error: %v", err)
// 	}
// }
// func TestFindLoginFromToken(t *testing.T) {
// 	validTokenString := "your_valid_jwt_token_here"
// 	invalidTokenString := "invalid_token"
// 	key := []byte("your_secret_key")

// 	t.Run("Valid Token", func(t *testing.T) {
// 		login, err := FindLoginFromToken(validTokenString, key)
// 		if err != nil {
// 			t.Errorf("Expected no error, but got: %v", err)
// 		}

// 		expectedLogin := "expected_login_value"
// 		if login != expectedLogin {
// 			t.Errorf("Expected login %s, but got %s", expectedLogin, login)
// 		}
// 	})

// 	t.Run("Invalid Token", func(t *testing.T) {
// 		_, err := FindLoginFromToken(invalidTokenString, key)
// 		if err == nil {
// 			t.Errorf("Expected error, but got nil")
// 		}
// 	})
// }

// func TestServer(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		switch r.URL.Path {
// 		case "/register":
// 			if r.Method != http.MethodPost {
// 				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 				return
// 			}
// 			username := r.FormValue("username")
// 			password := r.FormValue("password")
// 			fmt.Fprintf(w, "Registrating user with username: %s, password: %s", username, password)

// 		case "/login":
// 			// Handle login logic
// 		case "/checkjwt":
// 			// Handle JWT check logic
// 		case "/logout":
// 			// Handle logout logic
// 		default:
// 			http.Error(w, "Not Found", http.StatusNotFound)
// 		}
// 	}))
// 	defer ts.Close()

// 	res, err := http.Post(ts.URL+"/register", "application/json", strings.NewReader(`{"username": "testuser", "password": "testpassword"}`))
// 	if err != nil {
// 		t.Errorf("Failed to make POST request: %v", err)
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("Failed to read response body: %v", err)
// 	}

// 	fmt.Println(string(body))
// }

// func TestLoginHandler(t *testing.T) {
// 	//	db := NewDBConnection()

// 	// Регистрируем обработчики для каждого пути, используя замыкание для передачи db
// 	http.HandleFunc("/reg", RegApiRequests)
// 	http.HandleFunc("/login", LoginApiRequests)
// 	// http.HandleFunc("/password", PasswordHandler)

// 	// Запускаем сервер на порту :8080
// 	http.ListenAndServe(":8080", nil)
// 	fmt.Print("server is running")
// }

// func TestServer(t *testing.T) *OrchServer {
// 	var addr string
// 	var orch *Orchestrator

// 	mainCtx, stop := context.WithCancel(context.Background())

// 	rt := mux.NewRouter()
// 	http.Handle("/", rt)

// 	server := &http.Server{Addr: addr, Handler: rt, BaseContext: func(_ net.Listener) context.Context {
// 		return mainCtx
// 	}}

// 	orchServer := &OrchServer{httpServer: server, ctx: mainCtx, cancelFunc: stop, orch: orch}

// 	rt.HandleFunc("/add_expression", orchServer.addExpression)
// 	rt.HandleFunc("/get_expression", orchServer.getExpression)
// 	rt.HandleFunc("/set_expression_result", orchServer.setExpressionResult)
// 	rt.HandleFunc("/list_expressions", orchServer.listExpressions)
// 	rt.HandleFunc("/list_agents", orchServer.listExpressions)

// 	return orchServer
// }
