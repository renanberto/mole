package collectors

import "github.com/prometheus/client_golang/prometheus"

//OplogCollector - Collector metrics
type OplogCollector struct {
	oplogReadingCounter prometheus.Counter
}

//NewOplogCollector - Build the metrics collector
func NewOplogCollector() OplogCollector {
	collector := OplogCollector{
		oplogReadingCounter: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "oplog_reading_counter",
				Help: "Shows the amount of Oplog read from MongoDB",
			},
		),
	}

	prometheus.MustRegister(collector.oplogReadingCounter)
	return collector
}

//IncreasesReadingMetrics - Increases oplog reading metrics
func (o *OplogCollector) IncreasesReadingMetrics() {
	o.oplogReadingCounter.Inc()
}
