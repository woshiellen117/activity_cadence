package TestTrigger

import (
	"context"
	"go.uber.org/cadence/activity"
)

// Segment_produce workflow decider
func helloworldActivity(name string) (string, error) {
	logger := activity.GetLogger(context.Background())
	logger.Info("helloworld activity started")
	return "Hello " + name + "!", nil
}
