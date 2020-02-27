package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
)

const ApplicationName = "iw_mearth_3dtile_compiler"

func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope:          h.Scope,
		Logger:                h.Logger,
		DisableWorkflowWorker: true,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func main() {
	activity.RegisterWithOptions(iw_mearth_3dtile_compiler, activity.RegisterOptions{Name: "iw_mearth_3dtile_compiler"})
	var h common.SampleHelper
	h.SetupServiceConfig()
	startWorkers(&h)
	select {}
}
