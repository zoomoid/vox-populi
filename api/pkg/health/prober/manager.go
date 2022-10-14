package prober

import (
	"sync"
	"time"

	"github.com/go-logr/logr"
)

type ProbeManager struct {
	workers map[string]*worker

	workerLock sync.RWMutex

	start time.Time

	resultsManager ResultsManager

	logger *logr.Logger
}

func (m *ProbeManager) Add(pr prober, cfg *probeConfig) {
	m.workerLock.Lock()
	defer m.workerLock.Unlock()
	key := pr.Name()
	w := newWorker(m, key, m.resultsManager, pr, cfg)
	m.logger.V(3).Info("Created new worker for probe", "probe", key)
	m.workers[key] = w
	go w.run()
	m.logger.V(3).Info("Starting worker for probe in goroutine", "probe", key)
}

func (m *ProbeManager) Stop(pr prober) {
	m.workerLock.RLock()
	defer m.workerLock.RUnlock()

	key := pr.Name()

	if worker, ok := m.workers[key]; ok {
		m.logger.V(3).Info("stopping worker for probe", "probe", key)
		worker.stop()
	}
}
