package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type Expression struct {
	Id         string
	Expression string
	Result     string
	Status     string
}

type ExpressionAgent struct {
	address  string
	clientId string
	working  bool
	client   *http.Client
	wg       sync.WaitGroup
}

func NewExpressionAgent(address string, clientId string) *ExpressionAgent {
	var client *http.Client = &http.Client{}

	return &ExpressionAgent{
		address:  address,
		working:  false,
		client:   client,
		clientId: clientId,
	}
}

func (ea *ExpressionAgent) Start() {
	ea.working = true
	go ea.mainLoop()
	ea.wg.Add(1)
}

func (ea *ExpressionAgent) StopTimeout(timeout time.Duration) error {

	c := make(chan struct{})
	go func() {
		defer close(c)
		ea.wg.Wait()
	}()
	select {
	case <-c:
		return nil
	case <-time.After(timeout):

		return errors.New("exit by timeout")
	}
}

func (ea *ExpressionAgent) Stop() {
	ea.working = false
	ea.wg.Wait()
}

func (ea *ExpressionAgent) mainLoop() {
	defer ea.wg.Done()
	for ea.working {

		resp, err := ea.client.Post("http://localhost:8090/get_expression", "application/json", bytes.NewBuffer([]byte(ea.clientId)))

		if err != nil {
			time.Sleep(30 * time.Millisecond)
			continue
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		if len(bodyBytes) > 0 {

			var expr Expression
			err := json.Unmarshal(bodyBytes, &expr)
			if err == nil {
				doCalc(ea.clientId, expr, ea.client)

				time.Sleep(3000 * time.Millisecond)
			}
		}

		//		bodyString := string(bodyBytes)

		time.Sleep(2000 * time.Millisecond)
	}
}

func doCalc(id string, exp Expression, httpCl *http.Client) {

	result := "0"
	switch exp.Expression {
	case "+":
		result = fmt.Sprintf("%d", int(exp.Result[0])+int(exp.Result[1]))
	case "-":
		result = fmt.Sprintf("%d", int(exp.Result[0])-int(exp.Result[1]))
	case "*":
		result = fmt.Sprintf("%d", int(exp.Result[0])*int(exp.Result[1]))
	case "/":
		result = fmt.Sprintf("%d", int(exp.Result[0])/int(exp.Result[1]))
	}

	exp.Result = result
	exp.Status = id
	expBytes, _ := json.Marshal(exp)

	httpCl.Post("http://localhost:8090/set_expression_result", "application/json", bytes.NewBuffer(expBytes))
}

func main() {

	name := os.Args[1]

	fmt.Println("Starting calc name:" + name)

	var agent *ExpressionAgent = NewExpressionAgent("http://localhost:8090", name)
	agent.Start()

	fmt.Printf("Press Enter to stop")
	fmt.Fscanln(os.Stdin)
	fmt.Println("Stopping")

	agent.Stop()
	fmt.Println("Stopped")

}
