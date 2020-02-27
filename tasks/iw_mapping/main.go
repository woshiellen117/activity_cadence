package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
)

func init() {
	activity.RegisterWithOptions(iw_mapping, activity.RegisterOptions{
		Name: "iw_mapping",
	})
}

func main() {
	var h common.SampleHelper
	h.SetupServiceConfig()
	workerOptions := worker.Options{
		MetricsScope:          h.Scope,
		Logger:                h.Logger,
		DisableWorkflowWorker: true,
	}
	h.StartWorkers(h.Config.DomainName, "iw_mapping", workerOptions)
	select {}
}
