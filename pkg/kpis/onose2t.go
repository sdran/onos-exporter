// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpis

import (
	"github.com/onosproject/onos-lib-go/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
)

// Const definitions of kpis name and description associated with onose2t.
const (
	onosE2tConnectionsKPIName        = "connections"
	onosE2tConnectionsKPIDescription = "Indicates the number of e2t connections"
)

// Var definitions of e2t metrics builder and static labels.
// builder is used to create metrics in the PrometheusFormat.
var (
	staticLabelsE2t = map[string]string{"sdran": "e2t"}
	builder         = prom.NewBuilder("onos", "e2t", staticLabelsE2t)
)

// onosE2tConnections defines the common data that can be used
// to output the format of a KPI (e.g., PrometheusFormat).
type onosE2tConnections struct {
	name              string
	description       string
	Labels            []string
	LabelValues       []string
	NumberConnections float64
}

// PrometheusFormat implements the contract behavior of the kpis.KPI
// interface for onosE2tConnections.
func (c *onosE2tConnections) PrometheusFormat() (prometheus.Metric, error) {
	metricDesc := builder.NewMetricDesc(c.name, c.description, c.Labels, staticLabelsE2t)
	metric := builder.MustNewConstMetric(
		metricDesc,
		prometheus.GaugeValue,
		c.NumberConnections,
		c.LabelValues...,
	)

	return metric, nil
}

// OnosE2tConnections defines the factory implementation of a kpi
// onosE2tConnections having a well defined name and description.
func OnosE2tConnections() *onosE2tConnections {
	return &onosE2tConnections{
		name:        onosE2tConnectionsKPIName,
		description: onosE2tConnectionsKPIDescription,
	}
}
