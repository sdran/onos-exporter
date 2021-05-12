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

type onose2tCollectorPrometheus struct {
	E2tEndpoint string
}

func (c *onose2tCollectorPrometheus) Retrieve(ch chan<- prometheus.Metric) error {
	onose2tKPIs, err := collect.Onose2tCollector(c.E2tEndpoint)
	if err != nil {
		log.Errorf("onos-e2t collector error %s", err)
		return err
	}

	for _, e2tkpi := range onose2tKPIs {
		kpi, err := e2tkpi.PrometheusFormat()

		if err != nil {
			log.Errorf("onos-e2t kpi prometheus format error %s", err)
		} else {
			ch <- kpi
		}

	}

	return nil
}

func NewPrometheusExporter(config Config) prom.Exporter {
	exporter := prom.NewExporter(config.Path, config.Address)

	err := exporter.RegisterCollector("onos-e2t", &onose2tCollectorPrometheus{E2tEndpoint: config.E2tEndpoint})
	if err != nil {
		log.Errorf("error registering collector onos-e2t %s", err)
	}

	return exporter
}
