package prober

import "sync"

type Result int

type ResultsManager interface {
	Get(string) (Result, bool)

	Set(string, Result)

	Remove(string)

	Updates() <-chan Update
}

const (
	Unknown Result = iota - 1

	Success

	Failure
)

func (r Result) String() string {
	switch r {
	case Success:
		return "Success"
	case Failure:
		return "Failure"
	default:
		return "Unknown"
	}
}

// ToPrometheusType translates a Result to a form which is better understood by prometheus.
func (r Result) ToPrometheusType() float64 {
	switch r {
	case Success:
		return 0
	case Failure:
		return 1
	default:
		return -1
	}
}

type Update struct {
	ProbeType string
	Result    Result
}

type resultsManager struct {
	sync.RWMutex

	cache map[string]Result

	updates chan Update
}

var _ ResultsManager = &resultsManager{}

func NewResultsManager() ResultsManager {
	return &resultsManager{
		cache:   make(map[string]Result),
		updates: make(chan Update, 20),
	}
}

func (m *resultsManager) Get(probeType string) (Result, bool) {
	m.RLock()
	defer m.RUnlock()

	result, found := m.cache[probeType]

	return result, found
}

func (m *resultsManager) Set(probeType string, result Result) {
	if m.set(probeType, result) {
		m.updates <- Update{probeType, result}
	}
}

func (m *resultsManager) set(probeType string, result Result) bool {
	m.Lock()
	defer m.Unlock()

	prev, exists := m.cache[probeType]
	if !exists || prev != result {
		// is an actually new state
		m.cache[probeType] = result
		return true
	}
	return false
}

func (m *resultsManager) Remove(probeType string) {
	m.Lock()
	defer m.Unlock()
	delete(m.cache, probeType)
}

func (m *resultsManager) Updates() <-chan Update {
	return m.updates
}
