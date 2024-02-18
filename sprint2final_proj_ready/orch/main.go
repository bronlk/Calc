package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var orch *Orchestrator = NewOrchestrator()

	bytesRead, err := os.ReadFile("orch.cfg")

	if err == nil {
		var cfg OrchestratorCfg
		err := json.Unmarshal(bytesRead, &cfg)
		if err == nil {
			orch.ApplyCfg(cfg)
		}
	}

	s := NewServer("http://localhost:8090", orch)
	go s.Start()

	fmt.Printf("Press Enter to stop")
	fmt.Fscanln(os.Stdin)
	fmt.Println("Stopping")
	s.Stop()

	bytes, _ := json.Marshal(orch.GetCfg())

	prBytes := prettyprint(bytes)
	os.WriteFile("orch.cfg", prBytes, 0777)

	str := string(prettyprint(prBytes))
	fmt.Println(str)

	fmt.Println("Cfg saved")

	fmt.Println("Stopped")
}

// dont do this, see above edit
func prettyprint(b []byte) []byte {
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "  ")
	return out.Bytes()
}
