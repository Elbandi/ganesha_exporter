package main

import (
	"github.com/Gandi/ganesha_exporter/pkg/collectors"
	"github.com/Gandi/ganesha_exporter/pkg/dbus"
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	versioncollector "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promslog"
	"github.com/prometheus/common/promslog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"

	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"net/http"
	"os"
)

func main() {
	var (
		webConfig         = webflag.AddFlags(kingpin.CommandLine, ":9587")
		metricsPath       = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		gandi             = kingpin.Flag("gandi", "Activate Gandi specific fields").Default("false").Bool()
		exporterCollector = kingpin.Flag("collector.exports", "Activate exports collector").Default("true").Bool()
	)
	ec := collectors.NewExportsCollector()
	var clientCollector = kingpin.Flag("collector.clients", "Activate clients collector").Default("true").Bool()
	cc := collectors.NewClientsCollector()

	promslogConfig := &promslog.Config{}
	flag.AddFlags(kingpin.CommandLine, promslogConfig)
	kingpin.Version(version.Print("ctld_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promslog.New(promslogConfig)

	logger.Info("Starting ganesha_exporter", "version", version.Info())
	logger.Info("Build context", "context", version.BuildContext())

	dbus.Gandi = *gandi

	ec.InitDBus()
	cc.InitDBus()

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(versioncollector.NewCollector("ganesha_exporter"))
	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)
	if *exporterCollector {
		reg.MustRegister(ec)
	}
	if *clientCollector {
		reg.MustRegister(cc)
	}
	http.Handle(*metricsPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	if *metricsPath != "/" && *metricsPath != "" {
		landingConfig := web.LandingConfig{
			Name:        "ganesha_exporter",
			Description: "Prometheus Exporter for Ganesha nfs servers",
			Version:     version.Info(),
			Links: []web.LandingLinks{
				{
					Address: *metricsPath,
					Text:    "Metrics",
				},
			},
		}
		landingPage, err := web.NewLandingPage(landingConfig)
		if err != nil {
			logger.Error("Error creating landing page", "err", err)
			os.Exit(1)
		}
		http.Handle("/", landingPage)
	}

	srv := &http.Server{}
	if err := web.ListenAndServe(srv, webConfig, logger); err != nil {
		logger.Error("Error running HTTP server", "err", err)
		os.Exit(1)
	}
}
