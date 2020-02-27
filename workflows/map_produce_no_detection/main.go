package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
)

func init() {
	workflow.RegisterWithOptions(map_produce_no_detection, workflow.RegisterOptions{Name: "map_produce_no_detection"})
}
func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, "map_produce_no_detection", workerOptions)
}

func main() {
	var h common.SampleHelper
	h.SetupServiceConfig()
	startWorkers(&h)
	select {}
}
