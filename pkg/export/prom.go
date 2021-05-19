// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package export

import (
	"github.com/onosproject/onos-exporter/pkg/collect"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
)

var log = logging.GetLogger("export", "prom")

// onose2tCollectorPrometheus defines a prometheus collector
// for the Onose2tCollector. Only the E2tEndpoint service address
// field is required by Onose2tCollector, other fields
// might be added as needed by it.
type CollectorsPrometheus struct {
	collectors []collect.Collector
}

// Retrieve implements the method needed for a Collector interface
// in a prometheus exporter. It retrieves all the kpis from
// CollectorsPrometheus and pass them to the ch channel using the
// prometheus.Metric format.
// The function collect.KPIs performs the collection of each collector
// list of KPIs, and aggregates them in onosKPIs var.
func (c *CollectorsPrometheus) Retrieve(ch chan<- prometheus.Metric) error {
	onosKPIs, err := collect.KPIs(c.collectors)
	if err != nil {
		log.Errorf("onos collector error %s", err)
		return err
	}

	for _, kpi := range onosKPIs {
		promMetrics, err := kpi.PrometheusFormat()

		if err != nil {
			log.Errorf("onos kpi prometheus format error %s", err)
		} else {
			for _, m := range promMetrics {
				ch <- m
			}
		}

	}

	return nil
}

// Defines the set of collector used to extract KPIs for
// the prometheus exporter. Each collector implements the
// prom.Collector interface behavior via the method Collect.
func initCollectorsPrometheus(config Config) prom.Collector {
	e2tCollector := collect.Onose2tCollector{E2tServiceAddress: config.E2tEndpoint}
	e2subCollector := collect.Onose2subCollector{E2subServiceAddress: config.E2subEndpoint}
	xappPciCollector := collect.XappPciCollector{XappPciServiceAddress: config.XappPciEndpoint}
	xappKpimonCollector := collect.XappKpimonCollector{XappKpimonServiceAddress: config.XappKpimonEndpoint}

	return &CollectorsPrometheus{
		collectors: []collect.Collector{
			e2tCollector,
			e2subCollector,
			xappPciCollector,
			xappKpimonCollector,
		},
	}
}

// PrometheusExporter uses Config to create an instance of a
// Prometheus exporter, registering all its collectors, which must
// implement the interface method Retrieve.
func PrometheusExporter(config Config) prom.Exporter {
	exporter := prom.NewExporter(config.Path, config.Address)

	log.Info("Registering collector sdran")
	err := exporter.RegisterCollector("sdran", initCollectorsPrometheus(config))
	if err != nil {
		log.Errorf("error registering collector sdran %s", err)
	}

	return exporter
}
