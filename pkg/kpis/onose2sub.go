// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpis

import (
	"github.com/onosproject/onos-lib-go/pkg/prom"
	"github.com/prometheus/client_golang/prometheus"
)

// Var definitions of e2sub metrics builder and static labels.
// builder is used to create metrics in the PrometheusFormat.
var (
	staticLabelsE2sub = map[string]string{"sdran": "e2sub"}
	onose2tsubBuilder = prom.NewBuilder("onos", "e2sub", staticLabelsE2sub)
)

type E2Subscription struct {
	ID                  string
	Revision            string
	AppID               string
	ServiceModelName    string
	ServiceModelVersion string
	E2NodeID            string
	LifecycleStatus     string
}

// onosE2subs defines the common data that can be used
// to output the format of a KPI (e.g., PrometheusFormat).
// Subscriptions refer to each subscription of e2sub,
// given the annotations of that as defined by E2Subscription struct.
type onosE2subs struct {
	name          string
	description   string
	Labels        []string
	LabelValues   []string
	Subscriptions map[string]E2Subscription
}

// PrometheusFormat implements the contract behavior of the kpis.KPI
// interface for onosE2subs.
func (c *onosE2subs) PrometheusFormat() ([]prometheus.Metric, error) {
	metrics := []prometheus.Metric{}

	c.Labels = []string{"id", "revision", "appid", "service_model_name", "service_model_version", "e2nodeid", "lifecycle_status"}
	metricDesc := onose2tsubBuilder.NewMetricDesc(c.name, c.description, c.Labels, staticLabelsE2sub)

	for _, sub := range c.Subscriptions {
		metric := onose2tsubBuilder.MustNewConstMetric(
			metricDesc,
			prometheus.GaugeValue,
			1,
			sub.ID,
			sub.Revision,
			sub.AppID,
			sub.ServiceModelName,
			sub.ServiceModelVersion,
			sub.E2NodeID,
			sub.LifecycleStatus,
		)
		metrics = append(metrics, metric)
	}

	return metrics, nil
}
