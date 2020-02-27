package main

import (
	"encoding/json"
	"go.momenta.works/activity_cadence/conductor/model"
	"go.momenta.works/activity_cadence/conductor/transfer"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// Segment_produce workflow decider
func map_produce_no_detection(ctx workflow.Context, inputs []byte) ([]byte, error) {

	filePath := "map_produce_no_detection.json"
	dslWorkflow := transfer.TransferConductorTOCadence(filePath)

	var inputsJSON interface{}
	if err := json.Unmarshal(inputs, &inputsJSON); err != nil {
		return nil, err
	}
	inputMap := make(map[string]interface{})
	inputMap["input"] = inputsJSON

	bindings := make(map[string]interface{})
	bindings["workflow"] = inputMap

	ao := model.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 20,
		StartToCloseTimeout:    time.Minute * 20,
		HeartbeatTimeout:       time.Minute * 20,
	}
	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: ao.ScheduleToStartTimeout,
		StartToCloseTimeout:    ao.StartToCloseTimeout,
		HeartbeatTimeout:       ao.HeartbeatTimeout,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)
	logger := workflow.GetLogger(ctx)

	err := dslWorkflow.Root.Execute(ctx, ao, bindings)
	if err != nil {
		logger.Error("Segment_produce failed", zap.Error(err))
		return nil, err
	}
	logger.Info("Segment_produce completed")
	return nil, err
}
