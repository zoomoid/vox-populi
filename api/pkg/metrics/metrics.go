package metrics

type MetricsModule struct {
	registry *Registry
}

func New() *MetricsModule {
	return &MetricsModule{
		registry: &Registry{
			probes: map[string]*ProbeMetrics{
				"sqlite":    makeProbeMetrics("sqlite"),
				"memcached": makeProbeMetrics("memcached"),
			},
		},
	}
}

type Registry struct {
	probes map[string]*ProbeMetrics
}

func (module *MetricsModule) Get() *Registry {
	return module.registry
}

func (registry *Registry) Probes(key string) *ProbeMetrics {
	metrics, ok := registry.probes[key]
	if !ok {
		metrics = makeProbeMetrics(key)
		registry.probes[key] = metrics

	}
	return metrics
}
