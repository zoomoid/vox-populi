package prober

import (
	"time"
)

type worker struct {
	stopCh chan struct{}

	manualTriggerCh chan struct{}

	prober prober

	probeManager *ProbeManager

	resultsManager ResultsManager

	// The probe value during initial delay
	initialValue Result

	// last probe status
	lastResult Result

	// number of times the same result ooccured
	resultRun int

	// if set, skip probe
	onHold bool

	probeType string

	config *probeConfig
}

func newWorker(m *ProbeManager, probeType string, resultsManager ResultsManager, prober prober, cfg *probeConfig) *worker {
	w := &worker{
		stopCh:          make(chan struct{}, 1),
		manualTriggerCh: make(chan struct{}, 1),
		initialValue:    Unknown,
		probeManager:    m,
		prober:          prober,
		config:          cfg,
		probeType:       probeType,
		resultsManager:  resultsManager,
	}
	return w
}

func (w *worker) run() {
	probeTicker := time.NewTicker(w.config.period)

	defer func() {
		probeTicker.Stop()
	}()

probeLoop:
	for w.doProbe() {
		// Wait for next probe tick.
		select {
		case <-w.stopCh:
			break probeLoop
		case <-probeTicker.C:
		case <-w.manualTriggerCh:
			// continue
		}
	}
}

func (w *worker) stop() {
	select {
	case w.stopCh <- struct{}{}:
	default: // non-blocking
	}
}

func (w *worker) doProbe() (keepGoing bool) {
	defer func() { recover() }()

	// startTime := time.Now()

	if w.onHold {
		return true
	}

	result, err := w.prober.Probe()
	if err != nil {
		return true
	}

	switch result {

	}

	if w.lastResult == result {
		w.resultRun++
	} else {
		w.lastResult = result
		w.resultRun = 1
	}

	if (result == Failure && w.resultRun < int(w.config.failureThreshold)) ||
		(result == Success && w.resultRun < int(w.config.successThreshold)) {
		return true
	}

	w.resultsManager.Set(w.probeType, result)
	return true
}
