package model

type ConductorWorkflow struct {
	WorkflowName    string           `json:"name"`
	Tasks           []*ConductorTask `json:"tasks"`
	FailureWorkflow string           `json:"failureWorkflow"`
}

type ConductorTask struct {
	Name           string                      `json:"name"`
	TaskRefName    string                      `json:"taskReferenceName"`
	Input          map[string]interface{}      `json:"inputParameters"`
	Type           string                      `json:"type"`
	CaseValueParam string                      `json:"caseValueParam"`
	DecisionCases  map[string][]*ConductorTask `json:"decisionCases"`
}

//type DecisionCase struct{
//	YesTasks      []*ConductorTask  `json:"YES"`
//	NoTasks       []*ConductorTask  `json:"NO"`
//	TrueTasks     []*ConductorTask  `json:":true"`
//}
