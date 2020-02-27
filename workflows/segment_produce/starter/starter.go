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
	StartWorkflowSegmentProduce()
}

func StartWorkflowSegmentProduce() error {
	workflowInputs := make(map[string]interface{})
	//fileDirectory:=build.Default.GOPATH+"/activity_cadence/workflows/segment_produce/json_data/workflow_input.json"
	fileDirectory := "../json_data/workflow_input.json"
	workflowInputData, err := ioutil.ReadFile(fileDirectory)
	if err != nil {
		fmt.Printf("Fail to read workflow input file : %s\n", err)
		return err
	}
	if err := json.Unmarshal(workflowInputData, &workflowInputs); err != nil {
		fmt.Printf("Fail to unmarshal workflow input json: %s\n", err)
		return err
	}
	workflowInputJSON, err := json.Marshal(workflowInputs)
	if err != nil {
		log.Error("marshal error", err)
		return err
	}
	var h common.SampleHelper
	h.SetupServiceConfig()

	//filePath := "segment_produce.json"
	//workflow := transfer.TransferConductorTOCadence(filePath)
	//workflow.Variables=workflowInputs

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "segment_produce" + uuid.New(),
		TaskList:                        "segment_produce",
		ExecutionStartToCloseTimeout:    time.Minute * 20,
		DecisionTaskStartToCloseTimeout: time.Minute * 20,
	}
	h.StartWorkflow(workflowOptions, "segment_produce", workflowInputJSON)
	select {}
	return nil
}
