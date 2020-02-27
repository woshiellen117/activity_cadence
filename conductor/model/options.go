package model

import "time"

type ActivityOptions struct {
	ScheduleToStartTimeout time.Duration
	StartToCloseTimeout    time.Duration
	HeartbeatTimeout       time.Duration
}
