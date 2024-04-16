package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type OrchServer struct {
	httpServer *http.Server
	ctx        context.Context
	cancelFunc context.CancelFunc
	orch       *Orchestrator

	// orch
}

func (o *OrchServer) Start() {

	//TODO зачитываем из файла все что сохранили
	var err = o.httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func (o *OrchServer) Stop() {
	o.httpServer.Shutdown(o.ctx)
	//сохраняем в файл все что есть
}

func (o *OrchServer) addExpression(w http.ResponseWriter, r *http.Request) {

	resBody, _ := io.ReadAll(r.Body)
	expressionText := string(resBody)

	id := o.orch.Add(expressionText)
	w.Write([]byte(id))
}

// служит одновременно и пингом, каждые 10 миллсеккунд .
func (o *OrchServer) getExpression(w http.ResponseWriter, r *http.Request) {

	b, _ := io.ReadAll(r.Body)

	var calcId = string(b)

	success, expr := o.orch.GetExpression(calcId)

	if success {
		bytes, _ := json.Marshal(expr)
		w.Write(bytes)
	} else {
		w.Write([]byte(""))
	}
}

func (o *OrchServer) setExpressionResult(w http.ResponseWriter, r *http.Request) {

	var expr Expression
	json.NewDecoder(r.Body).Decode(&expr)
	o.orch.SetResult(expr)
}

func (o *OrchServer) listExpressions(w http.ResponseWriter, r *http.Request) {

	exprArr := o.orch.Expressions()

	for _, expr := range exprArr {
		bt, _ := json.Marshal(expr)
		w.Write(bt)
		w.Write([]byte("\n"))
	}
}

func (o *OrchServer) listAgents(w http.ResponseWriter, r *http.Request) {

	agents := o.orch.Agents()

	for _, agent := range agents {
		bt, _ := json.Marshal(agent)
		w.Write(bt)
		w.Write([]byte("\n"))
	}

	// bytes, _ := json.Marshal(agents)
	// w.Write(bytes)
}

func NewServer(addr string, orch *Orchestrator) *OrchServer {

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	rt := mux.NewRouter()
	http.Handle("/", rt)

	server := &http.Server{Addr: addr, Handler: rt,
		BaseContext: func(_ net.Listener) context.Context {
			return mainCtx
		}}

	orchServer := &OrchServer{httpServer: server, ctx: mainCtx, cancelFunc: stop, orch: orch}

	rt.HandleFunc("/add_expression", orchServer.addExpression)
	rt.HandleFunc("/get_expression", orchServer.getExpression)
	rt.HandleFunc("/set_expression_result", orchServer.setExpressionResult)
	rt.HandleFunc("/list_expressions", orchServer.listExpressions)
	rt.HandleFunc("/list_agents", orchServer.listAgents)
	//rt.HandleFunc("/list_calc", orchServer.list)

	return orchServer
}
