package main

import (
	"encoding/json"
	"fmt"
	"github.com/pborman/uuid"
	"github.com/prometheus/common/log"
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/client"
	"io/ioutil"
	"time"
)

func main() {

	workflowInputs := make(map[string]interface{})
	workflowInputData, err := ioutil.ReadFile("./workflows/map_produce_no_detection/json_data/workflow_input.json")
	if err != nil {
		fmt.Printf("Fail to read workflow input file : %s\n", err)
	}
	if err := json.Unmarshal(workflowInputData, &workflowInputs); err != nil {
		fmt.Printf("Fail to unmarshal workflow input json: %s\n", err)
	}
	workflowInputJSON, err := json.Marshal(workflowInputs)
	if err != nil {
		log.Error("marshal error", err)
	}
	var h common.SampleHelper
	h.SetupServiceConfig()

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "map_produce_no_detection" + uuid.New(),
		TaskList:                        "map_produce_no_detection",
		ExecutionStartToCloseTimeout:    time.Minute * 20,
		DecisionTaskStartToCloseTimeout: time.Minute * 20,
	}
	h.StartWorkflow(workflowOptions, "map_produce_no_detection", workflowInputJSON)
	select {}

}
