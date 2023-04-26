package prometheus_client

type MetricsModel struct {
	Labels    map[string]string
	Value     float64
	Timestamp int64
}
