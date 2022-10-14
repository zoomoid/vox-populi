package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ProbeMetrics struct {
	TimeoutsTotal               prometheus.Counter
	SuccessTotal                prometheus.Counter
	FailuresTotal               prometheus.Counter
	SuccessThresholdTotal       prometheus.Gauge
	FailureThresholdTotal       prometheus.Gauge
	TimeoutDurationSeconds      prometheus.Gauge
	PeriodDurationSeconds       prometheus.Gauge
	InitialDelayDurationSeconds prometheus.Gauge
}

func makeProbeMetrics(key string) *ProbeMetrics {
	timeoutsTotal := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "timeouts_total",
		Help:      "The total number of timeouts for the specific probe",
	})

	failuresTotal := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "failures_total",
		Help:      "The total number of failures for the specific probe",
	})

	successTotal := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "success_total",
		Help:      "The total number of successful tests for the specific probe",
	})

	successThresholdTotal := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "success_threshold_total",
		Help:      "The threshold of succesful tests that marks a probe as successful",
	})

	failureThresholdTotal := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "failure_threshold_total",
		Help:      "The threshold of succesful tests that marks a probe as failure",
	})

	timeoutsDurationSecond := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "timeout_duration_seconds",
		Help:      "The duration of the probe's timeout in seconds",
	})

	periodDurationSeconds := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "period_duration_seconds",
		Help:      "The period duration of the probe in seconds",
	})

	initialDelayDurationSeconds := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "probes",
		Subsystem: key,
		Name:      "initial_delay_duration_seconds",
		Help:      "The duration of the probe's initial delay in seconds",
	})

	return &ProbeMetrics{
		TimeoutsTotal:               timeoutsTotal,
		SuccessTotal:                successTotal,
		FailuresTotal:               failuresTotal,
		SuccessThresholdTotal:       successThresholdTotal,
		FailureThresholdTotal:       failureThresholdTotal,
		PeriodDurationSeconds:       periodDurationSeconds,
		TimeoutDurationSeconds:      timeoutsDurationSecond,
		InitialDelayDurationSeconds: initialDelayDurationSeconds,
	}
}
