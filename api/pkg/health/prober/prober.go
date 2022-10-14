package prober

type prober interface {
	Probe() (Result, error)

	Name() string
}
