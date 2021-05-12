// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpis

import (
	"github.com/onosproject/onos-lib-go/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	staticLabels = map[string]string{"sdran": "e2t"}
	builder      = prom.NewBuilder("onos", "e2t", staticLabels)
)

type OnosE2tConnections struct {
	name              string
	description       string
	Labels            []string
	LabelValues       []string
	NumberConnections float64
}

func (c *OnosE2tConnections) PrometheusFormat() (prometheus.Metric, error) {
	metricDesc := builder.NewMetricDesc(c.name, c.description, c.Labels, staticLabels)

	kpi := builder.MustNewConstMetric(
		metricDesc,
		prometheus.GaugeValue,
		c.NumberConnections,
		c.LabelValues...,
	)

	return kpi, nil

}

func NewOnosE2tConnections() *OnosE2tConnections {
	return &OnosE2tConnections{
		name:        "connections",
		description: "Indicates the number of e2t connections",
	}
}
