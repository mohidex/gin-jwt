package metricsutil

import "github.com/prometheus/client_golang/prometheus"

type PrometheusMetrics struct {
	RequestCounter      *prometheus.CounterVec
	UserCreationCounter prometheus.Counter
	RequestDuration     *prometheus.HistogramVec
}

func NewPrometheusMetrics() *PrometheusMetrics {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	userCreationCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "user_creation_total",
			Help: "Total number of user creations",
		},
	)

	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request durations in seconds",
			Buckets: []float64{0.1, 0.5, 1, 2, 5}, // Specify your buckets here
		},
		[]string{"method", "endpoint"},
	)

	prometheus.MustRegister(requestCounter, userCreationCounter, requestDuration)

	return &PrometheusMetrics{
		RequestCounter:      requestCounter,
		UserCreationCounter: userCreationCounter,
		RequestDuration:     requestDuration,
	}
}

func (pm *PrometheusMetrics) IncRequestCounter(method, endpoint, status string) {
	pm.RequestCounter.WithLabelValues(method, endpoint, status).Inc()
}

func (pm *PrometheusMetrics) IncUserCreationCounter() {
	pm.UserCreationCounter.Inc()
}

func (pm *PrometheusMetrics) ObserveRequestDuration(method, endpoint string, duration float64) {
	pm.RequestDuration.WithLabelValues(method, endpoint).Observe(duration)
}
