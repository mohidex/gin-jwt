package metricsutil

type Metrics interface {
	IncRequestCounter(method, endpoint, status string)
	IncUserCreationCounter()
	ObserveRequestDuration(method, endpoint string, duration float64)
}
