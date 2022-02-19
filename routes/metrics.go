package routes

// import (
//	"github.com/gin-gonic/gin"
//	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/promhttp"
//)
//
// func prometheusHandler() gin.HandlerFunc {
//	var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "cpu_temperature_celsius",
//		Help: "Current temperature of the CPU.",
//	})
//	prometheus.MustRegister(cpuTemp)
//	h := promhttp.Handler()
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
// }
//
// var metricsRoute routeFunc = func(Metrics *gin.RouterGroup) {
// 	Metrics.GET("", prometheusHandler())
// }
