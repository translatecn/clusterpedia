package server

import (
	"net/http"
	"time"

	"github.com/prometheus/exporter-toolkit/web"
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/klog/v2"

	_ "xxxxx/pkg/metrics"
	"xxxxx/pkg/pprof"
	"xxxxx/pkg/version"
)

type Config struct {
	Endpoint string

	TLSConfig           string
	DisableGZIPEncoding bool
}

func Run(config Config) {
	server := &http.Server{
		Handler:           buildMetricsServer(config),
		ReadHeaderTimeout: 6 * time.Second,
	}

	flags := &web.FlagConfig{
		WebListenAddresses: &[]string{config.Endpoint},
		WebSystemdSocket:   new(bool),
		WebConfigFile:      &config.TLSConfig,
	}

	klog.Info("Metrics Server is running...")
	_ = web.ListenAndServe(server, flags, Logger)
}

func buildMetricsServer(config Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/metrics", metrics.HandlerWithReset(
		legacyregistry.DefaultGatherer.(metrics.KubeRegistry),
		metrics.HandlerOpts{
			ErrorLog:           Logger,
			DisableCompression: config.DisableGZIPEncoding,
		}),
	)
	// add profiler
	pprof.RegisterProfileHandler(mux)
	// Add index
	landingConfig := web.LandingConfig{
		Name:        "clusterpedia clustersynchro manager",
		Description: "Self-metrics for clusterpedia clustersynchro manager",
		Version:     version.Get().String(),
		Links: []web.LandingLinks{
			{
				Text:    "Metrics",
				Address: "/metrics",
			},
		},
	}
	landingPage, err := web.NewLandingPage(landingConfig)
	if err != nil {
		klog.ErrorS(err, "failed to create landing page")
	}
	mux.Handle("/", landingPage)
	return mux
}
