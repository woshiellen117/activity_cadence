package activity_cadence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTransfer(t *testing.T) {
	workflowInput := make(map[string]interface{})
	workflowInputData, err := ioutil.ReadFile("./test_data/workflow_input.json")
	if err != nil {
		fmt.Sprintf("Fail to read workflow input file : %s", err)
	}
	if err := json.Unmarshal(workflowInputData, &workflowInput); err != nil {
		fmt.Sprintf("Fail to unmarshal workflow input json: %s", err)
	}

}
