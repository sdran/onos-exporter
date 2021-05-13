// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpis

import "github.com/prometheus/client_golang/prometheus"

// KPI interface defines the methods that format the behavior
// of a kpi. It includes that a kpi must provide those methods
// in order to support its content to be exported to a particular
// TSDB.
type KPI interface {
	PrometheusFormat() (prometheus.Metric, error)
}
