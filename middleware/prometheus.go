package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Go-Shop_requests_total",
			Help: "Total number of requests processed by the Go-Shop app ",
		},
		[]string{"path", "status"},
	)

	errCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Go-Shop_error_requests_total",
			Help: "Total number of error requeest processed by the Go-shop app",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(errCount)
}

func TrackMetrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		ctx.Next()
		status := ctx.Writer.Status()
		requestCount.WithLabelValues(path, http.StatusText(status)).Inc()
		if status >= 400 {
			errCount.WithLabelValues(path, http.StatusText(status)).Inc()
		}
	}
}
