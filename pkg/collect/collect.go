// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collect

import (
	"github.com/onosproject/onos-exporter/pkg/kpis"
)

var (
	onose2tConfig = InitConfig("onos-e2t")
)

func Onose2tCollector(e2tServiceAddress string) ([]kpis.KPI, error) {
	kpis := []kpis.KPI{}

	err := onose2tConfig.Set(map[string]string{addressKey: e2tServiceAddress})
	if err != nil {
		return kpis, err
	}

	conn, err := GetConnection(
		onose2tConfig.getAddress(),
		onose2tConfig.getCertPath(),
		onose2tConfig.getKeyPath(),
		onose2tConfig.noTLS(),
	)
	if err != nil {
		return kpis, err
	}
	defer conn.Close()

	e2tconnectionsKPI, err := onose2tListConnections(conn)
	if err != nil {
		return kpis, err
	}

	kpis = append(kpis, e2tconnectionsKPI)

	return kpis, nil
}
