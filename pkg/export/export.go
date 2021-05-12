// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package export

type Config struct {
	Address     string
	Path        string
	Mode        string
	CAPath      string
	KeyPath     string
	CertPath    string
	E2tEndpoint string
}

type exporter interface {
	Run() error
}

func NewExporter(cfg Config) exporter {
	switch cfg.Mode {
	case "prometheus":
		log.Info("Creating prometheus exporter")
		return NewPrometheusExporter(cfg)
	default:
		return NewPrometheusExporter(cfg)
	}
}
