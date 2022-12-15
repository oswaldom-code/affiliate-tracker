package infrastructure

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	metrics "github.com/penglongli/gin-metrics/ginmetrics"

	"github.com/oswaldom-code/affiliate-tracker/pkg/config"
	"github.com/oswaldom-code/affiliate-tracker/src/adapters/http/rest/handlers"
)

var basicAuthorizationMiddleware MiddlewareFunc = func(c *gin.Context) {
	// get token from header
	token := c.GetHeader("Authorization")
	// validate token
	if token != "Basic "+base64.StdEncoding.EncodeToString([]byte(config.GetAuthenticationKey().Secret)) {
		// response unauthorized status code
		c.Redirect(http.StatusFound, "/authorization")
	}
}

func setMetrics(router *gin.Engine) {
	// get global Monitor object
	monitor := metrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	monitor.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	monitor.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	monitor.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	monitor.Use(router)
}

// NewGinServer creates a new Gin server.
func NewGinServer(handler ServerInterface) *gin.Engine {
	// get configuration
	serverConfig := config.GetServerConfig()
	// validate parameters configuration
	if ok := serverConfig.Validate(); ok != nil {

		panic(ok)
	}

	// set gin mode (debug or release)
	gin.SetMode(serverConfig.Mode)
	// if debug mode is release write the logs to a file
	if serverConfig.Mode == "release" {
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()
		// Logging to a file.  // TODO: add current date to log file name
		f, ok := os.Create("log/error.log")
		if ok != nil {
			panic(ok)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// create routes
	router := gin.Default()
	// set static files directory
	router.Static("/static", serverConfig.Static)
	// set metrics
	setMetrics(router)
	// set middleware
	ginServerOptions := GinServerOptions{BaseURL: "/"}
	ginServerOptions.Middlewares = append(ginServerOptions.Middlewares, basicAuthorizationMiddleware)
	// register Handler, router and middleware to gin
	RegisterHandlersWithOptions(router, handler, ginServerOptions)
	return router
}

func loadHandlers() *handlers.Handler {
	return handlers.NewRestHandler()
}

func NewServer() *gin.Engine {
	return NewGinServer(
		loadHandlers(),
	)
}
