package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http" // Путь к вашим структурам

	"github.com/gorilla/mux"
)

type OrchRepository struct {
	httpServer *http.Server
	ctx        context.Context
	cancelFunc context.CancelFunc
	orch       *Orchestrator
	fileName   string // Используйте вашу структуру Orchestrator
}

func NewOrchRepository(fileName string) *OrchRepository {
	return &OrchRepository{fileName: fileName}
}

func (o *OrchRepository) Start() {

	//TODO зачитываем из файла все что сохранили
	var err = o.httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func (o *OrchRepository) Stop() {
	o.httpServer.Shutdown(o.ctx)
	//сохраняем в файл все что есть
}

func (o *OrchRepository) addExpression(w http.ResponseWriter, r *http.Request) {

	o.orch.SaveToDatabase(expressionText)
	return
}

// служит одновременно и пингом, каждые 10 миллсеккунд .
func (o *OrchRepository) getExpression(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	calcId := string(b)

	o.orch.GetExpression(calcId)
	return
}

func (o *OrchRepository) setExpressionResult(w http.ResponseWriter, r *http.Request) {

	var exp Expression
	err := json.NewDecoder(r.Body).Decode(&exp)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	o.orch.SetResult(exp)
	return
}

func (o *OrchRepository) listExpressions(w http.ResponseWriter, r *http.Request) {
	o.orch.PrintExpressions()
	return
}

func NewServer(addr string, orch *Orchestrator) *OrchRepository {
	mainCtx, stop := context.WithCancel(context.Background())

	rt := mux.NewRouter()
	http.Handle("/", rt)

	server := &http.Server{Addr: addr, Handler: rt, BaseContext: func(_ net.Listener) context.Context {
		return mainCtx
	}}

	orchServer := &OrchRepository{httpServer: server, ctx: mainCtx, cancelFunc: stop, orch: orch}

	rt.HandleFunc("/add_expression", orchServer.addExpression)
	rt.HandleFunc("/get_expression", orchServer.getExpression)
	rt.HandleFunc("/set_expression_result", orchServer.setExpressionResult)
	rt.HandleFunc("/list_expressions", orchServer.listExpressions)
	rt.HandleFunc("/list_agents", orchServer.listExpressions)

	return orchServer
}
