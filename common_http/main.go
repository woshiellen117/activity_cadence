package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
)

func init() {
	activity.RegisterWithOptions(CommonHttp, activity.RegisterOptions{
		Name: "common_http",
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
	h.StartWorkers(h.Config.DomainName, "common_http", workerOptions)
	select {}
}
