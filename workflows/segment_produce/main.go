package main

import (
	"github.com/uber-common/cadence-samples/cmd/samples/common"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
)

func init() {
	workflow.RegisterWithOptions(Segment_produce, workflow.RegisterOptions{Name: "segment_produce"})
}
func startWorkers(h *common.SampleHelper) {
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, "segment_produce", workerOptions)
}

func main() {
	var h common.SampleHelper
	h.SetupServiceConfig()
	startWorkers(&h)
	select {}
}
