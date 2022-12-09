package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	generalRequestsTotal           = prometheus.NewDesc("icap_general_requests_total", "General Statistics requests.", []string{}, nil)
	generalFailedRequestsTotal     = prometheus.NewDesc("icap_general_failed_requests_total", "General Statistics failed requests.", []string{}, nil)
	virusScanRequestsRespModsTotal = prometheus.NewDesc("icap_virus_scan_requests_total", "Service virus_scan Statistics RESPMODS requests", []string{}, nil)
	virusScanRequestsTotal         = prometheus.NewDesc("icap_virus_scan_requests_total", "Service virus_scan Statistics Requests scanned.", []string{}, nil)
	virusScanVirusFoundTotal       = prometheus.NewDesc("icap_virus_scan_virus_found_total", "Service virus_scan Statistics Viruses found.", []string{}, nil)
	virusScanFailureTotal          = prometheus.NewDesc("icap_virus_scan_failed_total", "Service virus_scan Statistics Scan failures.", []string{}, nil)
)

type icapCollector struct {
	config *Config
}

func (mc icapCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(mc, ch)
}

func (mc icapCollector) Collect(ch chan<- prometheus.Metric) {
	stats, err := getStats(mc.config)
	if err != nil {
		log.Printf("getStats error: %s", err.Error())
		stats = &Stats{}
	}
	ch <- prometheus.MustNewConstMetric(generalRequestsTotal, prometheus.CounterValue, stats.generalRequestsTotal)
	ch <- prometheus.MustNewConstMetric(generalFailedRequestsTotal, prometheus.CounterValue, stats.generalFailedRequestsTotal)
	ch <- prometheus.MustNewConstMetric(virusScanRequestsRespModsTotal, prometheus.CounterValue, stats.virusScanRequestsRespModsTotal)
	ch <- prometheus.MustNewConstMetric(virusScanRequestsTotal, prometheus.CounterValue, stats.virusScanRequestsScannedTotal)
	ch <- prometheus.MustNewConstMetric(virusScanVirusFoundTotal, prometheus.CounterValue, stats.virusScanVirusFoundTotal)
	ch <- prometheus.MustNewConstMetric(virusScanFailureTotal, prometheus.CounterValue, stats.virusScanFailureTotal)
}

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal("Unable to setup config")
	}

	prometheus.MustRegister(&icapCollector{config: config})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		_, err := execCmd(config)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}

	})

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening on port ", config.ServicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.ServicePort), nil))
}
