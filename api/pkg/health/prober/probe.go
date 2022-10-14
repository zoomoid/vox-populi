package prober

import "time"

type probeConfig struct {
	initialDelay     time.Duration
	period           time.Duration
	timeout          time.Duration
	successThreshold int64
	failureThreshold int64
}

type probeCondition string

var (
	ProbeConditionUnknown   probeCondition = "Unknown"
	ProbeConditionUnhealthy probeCondition = "Unhealthy"
	ProbeConditionHealthy   probeCondition = "Healthy"
)
