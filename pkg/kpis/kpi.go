// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpis

import "github.com/prometheus/client_golang/prometheus"

type KPI interface {
	PrometheusFormat() (prometheus.Metric, error)
}
