package TestTrigger

import (
	"context"
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestTrigger(t *testing.T) {
	var h common.SampleHelper
	h.SetupServiceConfig()
	workflowOptions := client.StartWorkflowOptions{
		TaskList:                        "segment_produce",
		ExecutionStartToCloseTimeout:    time.Minute * 20,
		DecisionTaskStartToCloseTimeout: time.Minute * 20,
	}
	workflowClient, err := h.Builder.BuildCadenceClient()
	if err != nil {
		h.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}

	we, err := workflowClient.SignalWithStartWorkflow(context.Background(), "segment_producefef83c93-20e8-41b1-bff4-0360d9c274a6", "wait_iw_fm_manu_lane", "SOME_VALUE", workflowOptions, "segment_produce", "")

	//we, err := workflowClient.SignalWithStartWorkflow(context.Background(), "segment_producefef83c93-20e8-41b1-bff4-0360d9c274a6", "wait_iw_fm_manu_delete_line", "SOME_VALUE", workflowOptions,"segment_produce","")
	//we, err := workflowClient.SignalWithStartWorkflow(context.Background(), "segment_producefef83c93-20e8-41b1-bff4-0360d9c274a6", "wait_manu_check_coverages", "SOME_VALUE", workflowOptions,"segment_produce","")
	if err != nil {
		h.Logger.Error("Failed to signal with start workflow", zap.Error(err))
		panic("Failed to signal with start workflow.")

	} else {
		h.Logger.Info("Signaled and started Workflow", zap.String("WorkflowID", we.ID), zap.String("RunID", we.RunID))
	}
	//h.SignalWithStartWorkflowWithCtx(context.Background(),"myDSLd94f9f05-8ffd-4eab-90ca-b61814884865","wait_iw_pole_manu_check_dup","SOME_VALUE",workflowOptions,"segment_produce")
}
