package main

func main() {

	//var orch *Orchestrator = NewOrchestrator()
	var databasePath = "../sqlite_db/sqlite.db"
	InitDB(databasePath)
	var userRepo *UserRepository = NewUserRepository(databasePath)
	var manager *UserManager = NewUserManager(userRepo)
	var userController *UserController = NewUserController(manager)
	var server *WebServer = NewWebServer(":8080", userController)
	server.Start()
	// var init *cUser

	// bytesRead, err := os.ReadFile("orch.cfg")

	// if err == nil {
	// 	var cfg OrchestratorCfg
	// 	err := json.Unmarshal(bytesRead, &cfg)
	// 	if err == nil {
	// 		orch.ApplyCfg(cfg)
	// 	}
	// }

	// s := NewServer("localhost:8080", orch)
	// go s.Start()

	// fmt.Printf("Press Enter to stop")
	// fmt.Fscanln(os.Stdin)
	// fmt.Println("Stopping")
	// s.Stop()

	// 	bytes, _ := json.Marshal(orch.GetCfg())

	// 	prBytes := prettyprint(bytes)
	// 	os.WriteFile("orch.cfg", prBytes, 0777)

	// 	str := string(prettyprint(prBytes))
	// 	fmt.Println(str)

	// 	fmt.Println("Cfg saved")

	// 	fmt.Println("Stopped")
	// }

	// dont do this, see above edit
	//
	//	func prettyprint(b []byte) []byte {
	//		var out bytes.Buffer
	//		_ = json.Indent(&out, b, "", "  ")
	//		return out.Bytes()
	//	}
}
