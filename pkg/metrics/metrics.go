// Package metrics provides Prometheus metrics for the futuapi4go SDK.
//
// Usage:
//
//	import "github.com/shing1211/futuapi4go/pkg/metrics"
//
//	// Before creating client
//	metrics.Init() // Start Prometheus server on :9090
//
//	// Or register custom handler
//	http.Handle("/metrics", metrics.Handler())
//
//	// Metrics are automatically collected
//	cli := client.New()
//
// Metrics tracked:
// - futuapi_connections_total (gauge)
// - futuapi_reconnects_total (counter)
// - futuapi_api_calls_total (counter)
// - futuapi_api_latency_seconds (histogram)
// - futuapi_push_messages_total (counter)
package metrics

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	initialized bool
	initMu      sync.Mutex

	connectionGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "futuapi_connections_total",
		Help: "Total number of active connections",
	}, []string{"type"})

	reconnectCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "futuapi_reconnects_total",
		Help: "Total number of reconnection attempts",
	}, []string{"reason"})

	apiCallsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "futuapi_api_calls_total",
		Help: "Total number of API calls",
	}, []string{"proto", "status"})

	apiLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "futuapi_api_latency_seconds",
		Help:    "API call latency in seconds",
		Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1},
	}, []string{"proto"})

	pushMessages = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "futuapi_push_messages_total",
		Help: "Total number of push messages received",
	}, []string{"type"})

	openDUp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "futuapi_opend_up",
		Help: "OpenD connection status (1=up, 0=down)",
	})

	lastConnectTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "futuapi_last_connect_timestamp",
		Help: "Timestamp of last successful connection",
	})

	apiErrorsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "futuapi_api_errors_total",
		Help: "Total number of API errors",
	}, []string{"proto", "error_code"})
)

func Init() {
	initMu.Lock()
	defer initMu.Unlock()
	if initialized {
		return
	}
	initialized = true
}

func InitWithServer(addr string) {
	Init()
	go func() {
		http.Handle("/metrics", Handler())
		http.ListenAndServe(addr, nil)
	}()
}

func Handler() http.Handler {
	return promhttp.Handler()
}

func RecordConnection(t string) {
	connectionGauge.WithLabelValues(t).Inc()
	lastConnectTime.Set(float64(time.Now().Unix()))
}

func RecordDisconnect(t string) {
	connectionGauge.WithLabelValues(t).Dec()
}

func RecordReconnect(reason string) {
	reconnectCounter.WithLabelValues(reason).Inc()
}

func RecordAPICall(proto, status string, latency time.Duration) {
	apiCallsTotal.WithLabelValues(proto, status).Inc()
	apiLatency.WithLabelValues(proto).Observe(latency.Seconds())
}

func RecordPushMessage(msgType string) {
	pushMessages.WithLabelValues(msgType).Inc()
}

func RecordOpenDUp(up bool) {
	val := 0.0
	if up {
		val = 1.0
	}
	openDUp.Set(val)
}

type APICallTracker struct {
	start    time.Time
	protoID  string
	finished bool
}

func StartAPITracking(protoID string) *APICallTracker {
	return &APICallTracker{
		start:   time.Now(),
		protoID: protoID,
	}
}

func (t *APICallTracker) End(success bool) {
	if t.finished {
		return
	}
	t.finished = true
	status := "success"
	if !success {
		status = "error"
	}
	latency := time.Since(t.start)
	RecordAPICall(t.protoID, status, latency)
}