package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
)

const ApplicationName = "iw_mesh"

func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope:          h.Scope,
		Logger:                h.Logger,
		DisableWorkflowWorker: true,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func main() {
	activity.RegisterWithOptions(iw_mesh, activity.RegisterOptions{Name: "iw_mesh"})
	var h common.SampleHelper
	h.SetupServiceConfig()
	startWorkers(&h)
	select {}
}
