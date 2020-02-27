package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
)

const ApplicationName = "iw_pole_auto_check_dup"

func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope:          h.Scope,
		Logger:                h.Logger,
		DisableWorkflowWorker: true,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func main() {
	activity.RegisterWithOptions(iw_pole_auto_check_dup, activity.RegisterOptions{Name: "iw_pole_auto_check_dup"})
	var h common.SampleHelper
	h.SetupServiceConfig()
	startWorkers(&h)
	select {}
}
