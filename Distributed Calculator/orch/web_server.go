package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type WebServer struct {
	httpServer *http.Server
	ctx        context.Context
	cancelFunc context.CancelFunc
	orch       *UserController

	// orch
}

func (o *WebServer) Start() {

	//TODO зачитываем из файла все что сохранили
	var err = o.httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func (o *WebServer) Stop() {
	o.httpServer.Shutdown(o.ctx)
	//сохраняем в файл все что есть
}

func NewWebServer(address string, userController *UserController, orchController *OrchController) *WebServer {

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	rt := mux.NewRouter()
	http.Handle("/", rt)

	server := &http.Server{Addr: address, Handler: rt,
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		}}

	orchServer := &WebServer{httpServer: server, ctx: mainCtx, cancelFunc: stop, orch: userController}

	rt.HandleFunc("/register", userController.RegisterUser)
	rt.HandleFunc("/login", userController.LoginApiRequest)
	rt.HandleFunc("/logout", userController.LogoutApiRequest)
	//rt.HandleFunc("/list_agents", orchServer.listAgents)
	//rt.HandleFunc("/list_calc", orchServer.list)
	rt.HandleFunc("/save_expression", orchController.AddExpressionByApi)
	rt.HandleFunc("/get_exoression", orchController.GetExpressionByApi)
	rt.HandleFunc("/list_exoression", orchController.PrintExpressionByApi)

	return orchServer
}
