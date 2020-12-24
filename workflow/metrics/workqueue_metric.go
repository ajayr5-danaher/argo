package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	WorkersBusyMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: argoNamespace,
			// TODO Subsystem:   workflowsSubsystem,
			// TODO should be `workers_busy` with an 's'
			Name: "worker_busy",
			Help: "Number of workers currently busy. https://argoproj.github.io/argo/fields/#workers_busy",
		},
		[]string{"queue_name"},
	)
)