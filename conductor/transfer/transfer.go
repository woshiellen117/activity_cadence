package transfer

import (
	"encoding/json"
	"fmt"
	"go.momenta.works/activity_cadence/conductor/cadence"
	"go.momenta.works/activity_cadence/conductor/model"
	"io/ioutil"
)

// LoadConductorWorkflow 加载conductor的json文件
func LoadConductorWorkflow(filePath string) (model.ConductorWorkflow, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Fail to read file : %s", err)
	}
	var workflow model.ConductorWorkflow
	if err := json.Unmarshal(data, &workflow); err != nil {
		fmt.Printf("Fail to unmarshal json file %s", err)
		return model.ConductorWorkflow{}, err
	}
	return workflow, nil
}

func TransferConductorTOCadence(filePath string) cadence.Workflow {

	conductorWorkflow, err := LoadConductorWorkflow("./conductor/transfer/json_data/" + filePath)
	if err != nil {
		fmt.Sprintf("Fail to Load Conductor workflow :%s", err)
	}
	workflow, err := buildCadence(conductorWorkflow)
	return workflow

}

func buildCadence(conductorWorkflow model.ConductorWorkflow) (cadence.Workflow, error) {
	statement, err := buildStatement(conductorWorkflow.Tasks)
	if err != nil {
		return cadence.Workflow{}, err
	}
	workflow := cadence.Workflow{
		Root: *statement,
	}
	return workflow, nil
}

func buildStatement(tasks []*model.ConductorTask) (*cadence.Statement, error) {
	var statements []*cadence.Statement
	for _, task := range tasks {
		if task.Type == "DECISION" {
			branches := make(map[string]cadence.Statement)
			for k, v := range task.DecisionCases {
				statement, err := buildStatement(v)
				if err != nil {
					return nil, err
				}
				branches[k] = *statement
			}
			choice := &cadence.Choice{
				Name:        task.Name,
				TaskRefName: task.TaskRefName,
				Arguments:   task.Input,
				Keybase:     task.CaseValueParam,
				Branches:    branches,
			}
			statement := &cadence.Statement{
				Activity: nil,
				Sequence: nil,
				Parallel: nil,
				Choice:   choice,
			}
			statements = append(statements, statement)
		} else {
			activity := &cadence.ActivityInvocation{
				Name:        task.Name,
				TaskRefName: task.TaskRefName,
				Arguments:   task.Input,
				Result:      "",
			}
			statement := &cadence.Statement{
				Activity: activity,
				Sequence: nil,
				Parallel: nil,
				Choice:   nil,
			}
			statements = append(statements, statement)
		}
	}
	sequence := &cadence.Sequence{Elements: statements}
	outerStatement := &cadence.Statement{
		Activity: nil,
		Sequence: sequence,
		Parallel: nil,
		Choice:   nil,
	}
	return outerStatement, nil
}
